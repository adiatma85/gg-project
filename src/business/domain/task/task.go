package task

import (
	"github.com/adiatma85/own-go-sdk/log"
	"github.com/adiatma85/own-go-sdk/parser"
	"github.com/adiatma85/own-go-sdk/redis"
	"github.com/adiatma85/own-go-sdk/sql"
)

type Interface interface {
	// Create(ctx context.Context)
	// Get()
	// GetList()
	// Update()
}
type InitParam struct {
	Log   log.Interface
	Db    sql.Interface
	Json  parser.JSONInterface
	Redis redis.Interface
}

type task struct {
	log   log.Interface
	db    sql.Interface
	json  parser.JSONInterface
	redis redis.Interface
}

func Init(param InitParam) Interface {
	t := &task{
		log:   param.Log,
		db:    param.Db,
		json:  param.Json,
		redis: param.Redis,
	}

	return t
}
