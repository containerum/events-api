package db

import (
	"time"

	"github.com/containerum/kube-client/pkg/model"
	"github.com/globalsign/mgo/bson"
)

func (mongo *MongoStorage) GetDeploymentEventsList(namespace, deployment string, startTime time.Time) ([]model.Event, error) {
	mongo.logger.Debugf("getting deployment events")
	var collection = mongo.db.C(DeploymentCollection)
	result := make([]model.Event, 0)
	if err := collection.Find(bson.M{
		"resourcenamespace": namespace,
		"resourcename":      deployment,
		"dateadded": bson.M{
			"$gte": startTime.Format(time.RFC3339),
		},
	}).All(&result); err != nil {
		mongo.logger.WithError(err).Errorf("unable to get deployment events")
		return nil, PipErr{error: err}.ToMongerr().NotFoundToNil().Extract()
	}
	return result, nil
}

func (mongo *MongoStorage) GetNamespaceDeploymentsEventsList(namespace string, startTime time.Time) ([]model.Event, error) {
	mongo.logger.Debugf("getting deployment events")
	var collection = mongo.db.C(DeploymentCollection)
	result := make([]model.Event, 0)
	if err := collection.Find(bson.M{
		"resourcenamespace": namespace,
		"dateadded": bson.M{
			"$gte": startTime.Format(time.RFC3339),
		},
	}).All(&result); err != nil {
		mongo.logger.WithError(err).Errorf("unable to get namespace deployments events")
		return nil, PipErr{error: err}.ToMongerr().NotFoundToNil().Extract()
	}
	return result, nil
}
