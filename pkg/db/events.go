package db

import (
	"time"

	"github.com/containerum/kube-client/pkg/model"
	"github.com/globalsign/mgo/bson"
)

func (mongo *MongoStorage) GetChangesList(namespace, resource, collectionName string, startTime time.Time) ([]model.Event, error) {
	mongo.logger.WithField("collection", collectionName).Debugf("getting changes")
	var collection = mongo.db.C(collectionName)
	result := make([]model.Event, 0)
	if err := collection.Find(bson.M{
		"resourcenamespace": namespace,
		"resourcename":      resource,
		"dateadded": bson.M{
			"$gte": startTime.Format(time.RFC3339),
		},
	}).All(&result); err != nil {
		mongo.logger.WithError(err).Errorf("unable to get changes")
		return nil, PipErr{error: err}.ToMongerr().NotFoundToNil().Extract()
	}
	return result, nil
}

func (mongo *MongoStorage) GetChangesInNamespaceList(namespace, collectionName string, startTime time.Time) ([]model.Event, error) {
	mongo.logger.Debugf("getting changes in namespace")
	var collection = mongo.db.C(collectionName)
	result := make([]model.Event, 0)
	if err := collection.Find(bson.M{
		"resourcenamespace": namespace,
		"dateadded": bson.M{
			"$gte": startTime.Format(time.RFC3339),
		},
	}).All(&result); err != nil {
		mongo.logger.WithError(err).Errorf("unable to get changes in namespace")
		return nil, PipErr{error: err}.ToMongerr().NotFoundToNil().Extract()
	}
	return result, nil
}
