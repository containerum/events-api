package db

import (
	"github.com/containerum/events-api/pkg/util/mongerr"
	"github.com/globalsign/mgo"
)

type PageInfo struct {
	PerPage        int
	Page           int
	DefaultPerPage int
}

func (pages PageInfo) Init() (limit, offset int) {
	if pages.PerPage <= 0 {
		if pages.DefaultPerPage > 0 {
			pages.PerPage = pages.DefaultPerPage
		} else {
			pages.PerPage = 100
		}
	}
	if pages.Page <= 0 {
		pages.Page = 0
	} else {
		pages.Page--
	}
	return pages.PerPage, pages.Page * pages.PerPage
}

func Paginate(query *mgo.Query, info *PageInfo) *mgo.Query {
	if info != nil {
		var limit, offset = info.Init()
		return query.Skip(offset).Limit(limit)
	}
	return query
}

type PipErr struct {
	error error
}

func (piperr PipErr) NotFoundToNil() PipErr {
	if piperr.error == mgo.ErrNotFound {
		return PipErr{}
	}
	return piperr
}

func (piperr PipErr) ToMongerr() PipErr {
	switch err := piperr.error.(type) {
	case *mgo.QueryError:
		return PipErr{error: mongerr.FromMongoErr(err)}
	default:
		return piperr
	}
}

func (piperr PipErr) Extract() error {
	return piperr.error
}

func (piperr PipErr) Apply(op func(error) error) PipErr {
	return PipErr{error: op(piperr.error)}
}
