package config

import (
	"flag"
	"fmt"
	"os"
)

var configs = map[string]Config{
	"default": {
		AmqpHost:     "localhost",
		AmqpUser:     "guest",
		AmqpPassword: "guest",
		HomePrefix:   "/Users/",
		UseLVE:       true,
	},

	"dev": {
		AmqpHost:     "zb.koding.com/kite",
		AmqpUser:     "guest",
		AmqpPassword: "s486auEkPzvUjYfeFTMQ",
		HomePrefix:   "/Users/",
		UseLVE:       true,
	},

	"dev-new": {
		AmqpHost:     "web0.dev.system.aws.koding.com:5672",
		AmqpUser:     "broker",
		AmqpPassword: "s486auEkPzvUjYfeFTMQ",
		HomePrefix:   "/Users/",
		UseLVE:       true,
	},

	"dev-new-web0": {
		AmqpHost:     "localhost:5672",
		AmqpUser:     "broker",
		AmqpPassword: "s486auEkPzvUjYfeFTMQ",
		HomePrefix:   "/Users/",
		UseLVE:       true,
	},

	"cl3-new": {
		AmqpHost:     "web0.dev.system.aws.koding.com:5672",
		AmqpUser:     "guest",
		AmqpPassword: "s486auEkPzvUjYfeFTMQ",
		HomePrefix:   "/Users/",
		UseLVE:       true,
	},

	"stage": {
		AmqpHost:     "web0.beta.system.aws.koding.com",
		AmqpUser:     "STAGE-sg46lU8J17UkVUq",
		AmqpPassword: "TV678S1WT221t1q",
		HomePrefix:   "/Users/",
		UseLVE:       true,
	},

	"prod": {
		AmqpHost:     "web0.beta.system.aws.koding.com",
		AmqpUser:     "prod-<component>",
		AmqpPassword: "Dtxym6fRJXx4GJz",
		HomePrefix:   "/Users/",
		UseLVE:       true,
	},

	"prod-new": {
		AmqpHost:     "web0.beta.system.aws.koding.com",
		AmqpUser:     "prod-<component>",
		AmqpPassword: "Dtxym6fRJXx4GJz",
		HomePrefix:   "/Users/",
		UseLVE:       true,
	},

	"local-go": {
		AmqpHost:     "localhost",
		AmqpUser:     "guest",
		AmqpPassword: "guest",
		HomePrefix:   "/home/",
	},

	"local": {
		AmqpHost:     "localhost",
		AmqpUser:     "guest",
		AmqpPassword: "guest",
		HomePrefix:   "/home/",
	},

	"websockets": {
		UseWebsockets: true,
		User:          "koding",
		HomePrefix:    "/home/",
	},
}

type Config struct {
	AmqpHost     string
	AmqpUser     string
	AmqpPassword string
	HomePrefix   string
	UseLVE       bool

	// for webterm's websockets mode
	UseWebsockets bool
	User          string
}

var Profile string
var Current Config

func init() {
	flag.StringVar(&Profile, "c", "", "Configuration profile")
}

func LoadConfig() {
	if Profile == "" {
		fmt.Println("Please specify a configuration profile (-c).")
		flag.PrintDefaults()
		os.Exit(1)
	}
	var ok bool
	Current, ok = configs[Profile]
	if !ok {
		fmt.Printf("Configuration not found: %v\n", Profile)
		os.Exit(1)
	}
}
