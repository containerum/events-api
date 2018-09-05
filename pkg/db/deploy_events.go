package db

import (
	"github.com/containerum/kube-client/pkg/model"
	"github.com/globalsign/mgo/bson"
)

func (mongo *MongoStorage) GetDeploymentEventsList(namespace, deployment string) ([]model.Event, error) {
	mongo.logger.Debugf("getting deployment events")
	var collection = mongo.db.C(DeploymentCollection)
	result := make([]model.Event, 0)
	if err := collection.Find(bson.M{
		"resourcenamespace": namespace,
		"resourcename":      deployment,
	}).All(&result); err != nil {
		mongo.logger.WithError(err).Errorf("unable to get deployment events")
		return nil, PipErr{error: err}.ToMongerr().NotFoundToNil().Extract()
	}
	return result, nil
}

func (mongo *MongoStorage) GetNamespaceDeploymentsEventsList(namespace string) ([]model.Event, error) {
	mongo.logger.Debugf("getting deployment events")
	var collection = mongo.db.C(DeploymentCollection)
	result := make([]model.Event, 0)
	if err := collection.Find(bson.M{
		"resourcenamespace": namespace,
	}).All(&result); err != nil {
		mongo.logger.WithError(err).Errorf("unable to get namespace deployments events")
		return nil, PipErr{error: err}.ToMongerr().NotFoundToNil().Extract()
	}
	return result, nil
}
