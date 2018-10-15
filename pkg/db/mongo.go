package db

import (
	"fmt"
	"time"

	"errors"

	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

const (
	localURL = "localhost:27017"
)

type MongoStorage struct {
	logger log.FieldLogger
	config MongoConfig
	closed bool
	db     *mgo.Database
}

func (mongo *MongoStorage) Close() (err error) {
	defer func() {
		switch rec := recover().(type) {
		case nil:
		case error:
			err = rec
		case fmt.Stringer:
			err = errors.New(rec.String())
		default:
			err = fmt.Errorf("%v", rec)
		}
	}()
	if mongo.closed {
		return fmt.Errorf("mongo stoarage already closed")
	}
	mongo.db.Session.Close()
	mongo.db = nil
	mongo.closed = true
	return nil
}

func (mongo *MongoStorage) IsClosed() bool {
	return mongo.closed
}

func NewMongo(config MongoConfig) (*MongoStorage, error) {
	if config.Logger == nil {
		var logger = log.StandardLogger()
		if config.Debug {
			logger.SetLevel(log.DebugLevel)
		}
		config.Logger = logger
	}
	if config.AppName == "" {
		config.AppName = "events-api"
	}
	config.FailFast = true
	config.Logger = config.Logger.WithField("app", config.AppName)
	if config.Debug {
		config.Logger.Debugf("running in debug mode")
	}
	config.Logger.Debugf("running mongo init")

	if config.Timeout <= 0 {
		config.Timeout = 10 * time.Second
	}
	config.Logger.Debugf("config timeout %v", config.Timeout)

	if len(config.Addrs) == 0 {
		config.Addrs = append(config.Addrs, localURL)
	}
	config.Logger.Debugf("addrs %v", config.Addrs)

	session, err := mgo.DialWithInfo(&config.DialInfo)
	if err != nil {
		config.Logger.WithError(err).Errorf("unable to connect to mongo")
		return nil, err
	}
	mgo.SetDebug(config.Debug)
	var db = session.DB(config.Database)
	if config.Username != "" || config.Password != "" {
		if err := db.Login(config.Username, config.Password); err != nil {
			return nil, err
		}
	}
	var storage = &MongoStorage{
		logger: config.Logger,
		config: config,
		db:     db,
	}
	return storage, nil
}
