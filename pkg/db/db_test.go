package db

import (
	"testing"

	"github.com/globalsign/mgo"
	"github.com/stretchr/testify/assert"
)

func TestDBConnection(t *testing.T) {
	dialInfo := mgo.DialInfo{Addrs: []string{"localhost:27017"}}
	cfg := MongoConfig{DialInfo: dialInfo}

	_, err := NewMongo(cfg)
	assert.Nil(t, err)
}
