package db

import (
	"time"

	"github.com/containerum/kube-client/pkg/model"
	"github.com/containerum/kube-events/pkg/storage/mongodb"
	"github.com/globalsign/mgo/bson"
)

func (mongo *MongoStorage) GetEventsList(namespace, resource string, resourcetype model.ResourceType, limit int, startTime time.Time) ([]model.Event, error) {
	mongo.logger.WithField("collection", mongodb.EventsCollection).Debugf("getting events")
	var collection = mongo.db.C(mongodb.EventsCollection)
	result := make([]model.Event, 0)
	if err := collection.Find(bson.M{
		"resourcenamespace": namespace,
		"resourcename":      resource,
		"resourcetype":      resourcetype,
		"dateadded": bson.M{
			"$gte": startTime,
		},
	}).Sort("-eventtime").Limit(limit).All(&result); err != nil {
		mongo.logger.WithError(err).Errorf("unable to get changes")
		return nil, PipErr{error: err}.ToMongerr().NotFoundToNil().Extract()
	}
	return result, nil
}

func (mongo *MongoStorage) GetEventsInNamespacesList(resourcetype model.ResourceType, limit int, startTime time.Time, namespaces ...string) ([]model.Event, error) {
	mongo.logger.WithField("collection", mongodb.EventsCollection).Debugf("getting events in namespace")
	var collection = mongo.db.C(mongodb.EventsCollection)
	result := make([]model.Event, 0)
	if err := collection.Find(bson.M{
		"resourcenamespace": bson.M{
			"$in": namespaces,
		},
		"resourcetype": resourcetype,
		"dateadded": bson.M{
			"$gte": startTime,
		},
	}).Sort("-eventtime").Limit(limit).All(&result); err != nil {
		mongo.logger.WithError(err).Errorf("unable to get events in namespace")
		return nil, PipErr{error: err}.ToMongerr().NotFoundToNil().Extract()
	}
	return result, nil
}

// Returns paginated event list (pageSize events per page).
// If pageN < 0, then uses default value 0.
// If pageSize < 1, then uses default value 32
// Results are sorted from newest to oldest.
func (mongo *MongoStorage) GetEventsInNamespacesListPaginated(pageN, pageSize int, namespaces ...string) ([]model.Event, error) {
	if pageN < 0 {
		pageN = 0
	}
	if pageSize < 1 {
		pageSize = 32
	}
	mongo.logger.WithField("collection", mongodb.EventsCollection).Debugf("getting events page %d(%d per page) in namespace", pageN, pageSize)
	var collection = mongo.db.C(mongodb.EventsCollection)
	result := make([]model.Event, 0)
	if err := collection.Find(bson.M{
		"resourcenamespace": bson.M{
			"$in": namespaces,
		},
	}).Sort("-eventtime").Skip(pageN * pageSize).Limit(pageSize).All(&result); err != nil {
		mongo.logger.WithError(err).Errorf("unable to get events page %d(%d per page) in namespace", pageN, pageSize)
		return nil, PipErr{error: err}.ToMongerr().NotFoundToNil().Extract()
	}
	return result, nil
}

func (mongo *MongoStorage) GetAllEventsList(resourcetype model.ResourceType, limit int, startTime time.Time) ([]model.Event, error) {
	mongo.logger.WithField("collection", mongodb.EventsCollection).Debugf("getting events in all namespaces")
	var collection = mongo.db.C(mongodb.EventsCollection)
	result := make([]model.Event, 0)
	if err := collection.Find(bson.M{
		"resourcetype": resourcetype,
		"dateadded": bson.M{
			"$gte": startTime,
		},
	}).Sort("-eventtime").Limit(limit).All(&result); err != nil {
		mongo.logger.WithError(err).Errorf("unable to get events in namespace")
		return nil, PipErr{error: err}.ToMongerr().NotFoundToNil().Extract()
	}
	return result, nil
}

func (mongo *MongoStorage) AddContainerumEvent(collectionName string, event model.Event) error {
	mongo.logger.WithField("collection", collectionName).Debugf("adding event")
	var collection = mongo.db.C(collectionName)
	if err := collection.Insert(event); err != nil {
		mongo.logger.WithError(err).Errorf("unable to add event")
		return PipErr{error: err}.ToMongerr().NotFoundToNil().Extract()
	}
	return nil
}
