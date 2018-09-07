package db

import (
	"time"

	"github.com/containerum/kube-client/pkg/model"
	"github.com/containerum/kube-events/pkg/storage/mongodb"
	"github.com/globalsign/mgo/bson"
)

func (mongo *MongoStorage) GetUsersEventsList(startTime time.Time) ([]model.Event, error) {
	mongo.logger.Debugf("getting users changes")
	var collection = mongo.db.C(mongodb.UserCollection)
	result := make([]model.Event, 0)
	if err := collection.Find(bson.M{
		"dateadded": bson.M{
			"$gte": startTime.Format(time.RFC3339),
		},
	}).All(&result); err != nil {
		mongo.logger.WithError(err).Errorf("unable to get users changes")
		return nil, PipErr{error: err}.ToMongerr().NotFoundToNil().Extract()
	}
	return result, nil
}
