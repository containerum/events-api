package main

import (
	"time"

	"github.com/containerum/events-api/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/en_US"
	"github.com/go-playground/universal-translator"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var flags = []cli.Flag{
	cli.BoolFlag{
		EnvVar: "CH_EVENTS_API_DEBUG",
		Name:   "debug",
		Usage:  "start the server in debug mode",
	},
	cli.BoolFlag{
		EnvVar: "CH_EVENTS_API_TEXTLOG",
		Name:   "textlog",
		Usage:  "output log in text format",
	},
	cli.StringFlag{
		EnvVar: "CH_EVENTS_API_PORT",
		Name:   "port",
		Value:  "1667",
		Usage:  "port for events-api server",
	},
	cli.BoolFlag{
		EnvVar: "CH_EVENTS_API_CORS",
		Name:   "cors",
		Usage:  "enable CORS",
	},
	cli.StringFlag{
		EnvVar: "CH_EVENTS_API_MONGO_DB",
		Name:   "mongo_db",
		Usage:  "MongoDB database name",
	},
	cli.StringFlag{
		EnvVar: "CH_EVENTS_API_MONGO_LOGIN",
		Name:   "mongo_login",
		Usage:  "MongoDB login",
	},
	cli.StringFlag{
		EnvVar: "CH_EVENTS_API_MONGO_PASSWORD",
		Name:   "mongo_password",
		Usage:  "MongoDB password",
	},
	cli.StringSliceFlag{
		EnvVar: "CH_EVENTS_API_MONGO_ADDR",
		Name:   "mongo_addr",
		Usage:  "MongoDB address",
	},
	cli.DurationFlag{
		EnvVar: "DB_REQUEST_PERIOD",
		Name:   "db_request_period",
		Usage:  "DB requests period (for websocket requests)",
		Value:  60 * time.Second,
	},
}

func setupLogs(c *cli.Context) {
	if c.Bool("debug") {
		gin.SetMode(gin.DebugMode)
		log.SetLevel(log.DebugLevel)
	} else {
		gin.SetMode(gin.ReleaseMode)
		log.SetLevel(log.InfoLevel)
	}

	if c.Bool("textlog") {
		log.SetFormatter(&log.TextFormatter{})
	} else {
		log.SetFormatter(&log.JSONFormatter{})
	}
}

func setupTranslator() *ut.UniversalTranslator {
	return ut.New(en.New(), en.New(), en_US.New())
}

func setupMongo(c *cli.Context) (*db.MongoStorage, error) {
	dialInfo := mgo.DialInfo{
		Username:  c.String("mongo_login"),
		Password:  c.String("mongo_password"),
		Addrs:     c.StringSlice("mongo_addr"),
		Database:  c.String("mongo_db"),
		Mechanism: "SCRAM-SHA-1",
	}
	cfg := db.MongoConfig{
		Logger:   log.WithField("component", "mongo"),
		Debug:    c.Bool("debug"),
		DialInfo: dialInfo,
	}
	return db.NewMongo(cfg)
}
