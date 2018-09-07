package db

import (
	"time"

	"github.com/containerum/kube-client/pkg/model"
	"github.com/globalsign/mgo/bson"
)

func (mongo *MongoStorage) GetEventsList(namespace, resource, resourcetype string, startTime time.Time) ([]model.Event, error) {
	mongo.logger.WithField("collection", EventsCollection).Debugf("getting events")
	var collection = mongo.db.C(EventsCollection)
	result := make([]model.Event, 0)
	if err := collection.Find(bson.M{
		"resourcenamespace": namespace,
		"resourcename":      resource,
		"resourcetype":      resourcetype,
		"dateadded": bson.M{
			"$gte": startTime.Format(time.RFC3339),
		},
	}).All(&result); err != nil {
		mongo.logger.WithError(err).Errorf("unable to get changes")
		return nil, PipErr{error: err}.ToMongerr().NotFoundToNil().Extract()
	}
	return result, nil
}

func (mongo *MongoStorage) GetEventsInNamespaceList(namespace, resourcetype string, startTime time.Time) ([]model.Event, error) {
	mongo.logger.Debugf("getting events in namespace")
	var collection = mongo.db.C(EventsCollection)
	result := make([]model.Event, 0)
	if err := collection.Find(bson.M{
		"resourcenamespace": namespace,
		"resourcetype":      resourcetype,
		"dateadded": bson.M{
			"$gte": startTime.Format(time.RFC3339),
		},
	}).All(&result); err != nil {
		mongo.logger.WithError(err).Errorf("unable to get events in namespace")
		return nil, PipErr{error: err}.ToMongerr().NotFoundToNil().Extract()
	}
	return result, nil
}