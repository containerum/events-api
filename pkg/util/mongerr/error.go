package mongerr

type Error struct {
	Message string `json:"message" bson:"message"`
	Code    int    `json:"code" bson:"code"`
}

func (err Error) Error() string {
	return err.Message
}
