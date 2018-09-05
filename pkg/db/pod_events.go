package db

import (
	"time"

	"github.com/containerum/kube-client/pkg/model"
	"github.com/globalsign/mgo/bson"
)

func (mongo *MongoStorage) GetPodEventsList(namespace, pod string, startTime time.Time) ([]model.Event, error) {
	mongo.logger.Debugf("getting pod events")
	var collection = mongo.db.C(PodEventsCollection)
	result := make([]model.Event, 0)
	if err := collection.Find(bson.M{
		"resourcenamespace": namespace,
		"resourcename":      pod,
		"dateadded": bson.M{
			"$gte": startTime.Format(time.RFC3339),
		},
	}).All(&result); err != nil {
		mongo.logger.WithError(err).Errorf("unable to get pod events")
		return nil, PipErr{error: err}.ToMongerr().NotFoundToNil().Extract()
	}
	return result, nil
}

func (mongo *MongoStorage) GetNamespacePodsEventsList(namespace string, startTime time.Time) ([]model.Event, error) {
	mongo.logger.Debugf("getting pods events")
	var collection = mongo.db.C(PodEventsCollection)
	result := make([]model.Event, 0)
	if err := collection.Find(bson.M{
		"resourcenamespace": namespace,
		"dateadded": bson.M{
			"$gte": startTime.Format(time.RFC3339),
		},
	}).All(&result); err != nil {
		mongo.logger.WithError(err).Errorf("unable to get namespace pods events")
		return nil, PipErr{error: err}.ToMongerr().NotFoundToNil().Extract()
	}
	return result, nil
}
