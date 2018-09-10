package db

import (
	"time"

	"github.com/containerum/kube-events/pkg/storage/mongodb"

	"github.com/containerum/kube-client/pkg/model"
	"github.com/globalsign/mgo/bson"
)

func (mongo *MongoStorage) GetNamespaceChangesList(namespace string, startTime time.Time) ([]model.Event, error) {
	mongo.logger.Debugf("getting namespace changes")
	var collection = mongo.db.C(mongodb.ResourceQuotasCollection)
	result := make([]model.Event, 0)
	if err := collection.Find(bson.M{
		"resourcename": namespace,
		"dateadded": bson.M{
			"$gte": startTime,
		},
	}).All(&result); err != nil {
		mongo.logger.WithError(err).Errorf("unable to get namespace changes")
		return nil, PipErr{error: err}.ToMongerr().NotFoundToNil().Extract()
	}
	return result, nil
}
