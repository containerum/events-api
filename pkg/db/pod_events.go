package db

import (
	"github.com/containerum/kube-client/pkg/model"
	"github.com/globalsign/mgo/bson"
)

func (mongo *MongoStorage) GetPodEventsList(namespace, pod string) ([]model.Event, error) {
	mongo.logger.Debugf("getting pod events")
	var collection = mongo.db.C(PodEventsCollection)
	result := make([]model.Event, 0)
	if err := collection.Find(bson.M{
		"resourcenamespace": namespace,
		"resourcename":      pod,
	}).All(&result); err != nil {
		mongo.logger.WithError(err).Errorf("unable to get pod events")
		return nil, PipErr{error: err}.ToMongerr().NotFoundToNil().Extract()
	}
	return result, nil
}

func (mongo *MongoStorage) GetNamespacePodsEventsList(namespace string) ([]model.Event, error) {
	mongo.logger.Debugf("getting pods events")
	var collection = mongo.db.C(PodEventsCollection)
	result := make([]model.Event, 0)
	if err := collection.Find(bson.M{
		"resourcenamespace": namespace,
	}).All(&result); err != nil {
		mongo.logger.WithError(err).Errorf("unable to get namespace pods events")
		return nil, PipErr{error: err}.ToMongerr().NotFoundToNil().Extract()
	}
	return result, nil
}
