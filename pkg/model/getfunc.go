package model

import (
	"time"

	"github.com/containerum/kube-client/pkg/model"
	"github.com/gin-gonic/gin"
)

type FuncParams struct {
	Params         gin.Params
	UserAdmin      bool
	UserNamespaces []string
	Limit          int
	StartTime      time.Time
	Page           int
	PageSize       int
}

type EventsFunc func(params FuncParams) (*model.EventsList, error)
