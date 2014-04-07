package channel

import (
	"net/http"
	"net/url"
	"socialapi/models"
	"socialapi/workers/api/modules/helpers"

	"github.com/jinzhu/gorm"
)

func validateChannelRequest(c *models.Channel) error {
	if c.GroupName == "" {
		return errors.New("Group name is not set")
	}

	if c.Name == "" {
		return errors.New("Channel name is not set")
	}

	if c.CreatorId == 0 {
		return errors.New("Creator id is not set")
	}

	return nil
}

func Create(u *url.URL, h http.Header, req *models.Channel) (int, http.Header, interface{}, error) {
	if err := req.Create(); err != nil {
		return helpers.NewBadRequestResponse(err)
	}

	return helpers.NewOKResponse(req)
}

func List(u *url.URL, h http.Header, _ interface{}) (int, http.Header, interface{}, error) {
	c := models.NewChannel()
	list, err := c.List(helpers.GetQuery(u))
	if err != nil {
		return helpers.NewBadRequestResponse(err)
	}

	return helpers.NewOKResponse(list)
}

func Delete(u *url.URL, h http.Header, req *models.Channel) (int, http.Header, interface{}, error) {
	id, err := helpers.GetURIInt64(u, "id")
	if err != nil {
		return helpers.NewBadRequestResponse(err)
	}

	req.Id = id

	if err := req.Delete(); err != nil {
		return helpers.NewBadRequestResponse(err)
	}
	// yes it is deleted but not removed completely from our system
	return helpers.NewDeletedResponse()
}

func Update(u *url.URL, h http.Header, req *models.Channel) (int, http.Header, interface{}, error) {
	id, err := helpers.GetURIInt64(u, "id")
	if err != nil {
		return helpers.NewBadRequestResponse(err)
	}
	req.Id = id

	if req.Id == 0 {
		return helpers.NewBadRequestResponse(err)
	}

	if err := req.Update(); err != nil {
		return helpers.NewBadRequestResponse(err)
	}

	return helpers.NewOKResponse(req)
}

func Get(u *url.URL, h http.Header, req *models.Channel) (int, http.Header, interface{}, error) {
	id, err := helpers.GetURIInt64(u, "id")
	if err != nil {
		return helpers.NewBadRequestResponse(err)
	}

	req.Id = id
	if err := req.Fetch(); err != nil {
		if err == gorm.RecordNotFound {
			return helpers.NewNotFoundResponse()
		}
		return helpers.NewBadRequestResponse(err)
	}

	return helpers.NewOKResponse(req)
}

func PostMessage(u *url.URL, h http.Header, req *models.Channel) (int, http.Header, interface{}, error) {
	// id, err := helpers.GetURIInt64(u, "id")
	// if err != nil {
	// 	return helpers.NewBadRequestResponse(err)
	// }

	// req.Id = id
	// // TODO - check if the user is member of the channel

	// if err := req.Fetch(); err != nil {
	// 	if err == gorm.RecordNotFound {
	// 		return helpers.NewNotFoundResponse()
	// 	}
	// 	return helpers.NewBadRequestResponse(err)
	// }

	return helpers.NewOKResponse(req)
}
