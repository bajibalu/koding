package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"koding/newkite/kodingkey"
	"koding/newkite/protocol"
	"koding/newkite/token"
	"net/http"
)

// everyone needs a place for home
func homeHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world - kontrol!\n")
}

// preparHandler first checks if the incoming POST request is a valid session.
// Every request made to kontrol should be in POST with protocol.KontrolQuery in
// their body.
func prepareHandler(fn func(w http.ResponseWriter, r *http.Request, q *query)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		q, err := readBrowserRequest(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("{\"err\":\"%s\"}\n", err), http.StatusBadRequest)
			return
		}

		err = q.Validate()
		if err != nil {
			http.Error(w, fmt.Sprintf("{\"err\":\"%s\"}\n", err), http.StatusBadRequest)
			return
		}

		log.Info("sessionID '%s' is validated as: %s", q.Authentication.Key, q.Username)
		fn(w, r, q)
	}
}

// we assume that the incoming JSON data is in form of protocol.KontrolQuery.
// Read and return a new protocol.KontrolQuery from the POST body if
// succesfull.
func readBrowserRequest(requestBody io.ReadCloser) (*query, error) {
	body, err := ioutil.ReadAll(requestBody)
	if err != nil {
		return nil, err
	}
	defer requestBody.Close()

	var req query
	err = json.Unmarshal(body, &req)
	if err != nil {
		return nil, err
	}

	return &req, nil
}

// searchForKites returns a list of kites that matches the variable matchKite
// It also generates a new one-way token that is used between the client and
// kite and appends it to each kite struct
func searchForKites(username, kitename string) ([]protocol.KiteWithToken, error) {
	kites := make([]protocol.KiteWithToken, 0)

	log.Info("searching for kite '%s'", kitename)

	for _, k := range storage.List() {
		if k.Username == username && k.Name == kitename {
			key, err := kodingkey.FromString(k.KodingKey)
			if err != nil {
				return nil, fmt.Errorf("Koding Key is invalid at Kite: %s", k.Name)
			}

			// username is from requester, key is from kite owner
			tokenString, err := token.NewToken(username, k.ID).EncryptString(key)
			if err != nil {
				return nil, errors.New("Server error: Cannot generate a token")
			}

			kwt := protocol.KiteWithToken{
				Kite:  k.Kite,
				Token: tokenString,
			}
			kites = append(kites, kwt)
		}
	}

	if len(kites) == 0 {
		return nil, fmt.Errorf("'%s' not available", kitename)
	}

	return kites, nil
}

// requestHandler sends as response a list of kites that matches kites in form
// of "username/kitename".
// Request comes from a web browser.
func requestHandler(w http.ResponseWriter, r *http.Request, q *query) {
	kites, err := searchForKites(q.Username, q.Kitename)
	if err != nil {
		errMsg := fmt.Sprintf("{\"err\":\"%s\"}\n", err)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	kitesJSON, err := json.Marshal(kites)
	if err != nil {
		http.Error(w, fmt.Sprintf("{\"err\":\"%s\"}\n", err), http.StatusBadRequest)
		return
	}

	w.Write([]byte(kitesJSON))
}
