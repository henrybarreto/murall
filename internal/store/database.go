package store

import "context"

type Database interface {
	OpenConnection(ctx *context.Context) interface{}
	CloseConnection(connection interface{}, ctx *context.Context) error
	SaveMsg(connection interface{}, ctx *context.Context, msg string) (interface{}, error)
	GetMsg(connection interface{}, ctx *context.Context) (interface{}, error)
}
