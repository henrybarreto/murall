package services

import (
	"time"

	"github.com/henrybarreto/murall/internal/store"
	"github.com/henrybarreto/murall/pkg/cacher"
	"golang.org/x/net/context"
)

var cache *cacher.Cache

func init() {
	cache = new(cacher.Cache)
}

func SaveMsg(msg string) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	database := new(store.DatabaseMongo)
	connection := database.OpensConnection(ctx)
	defer database.CloseConnection(connection, ctx)

	res, err := database.SaveMsg(connection, ctx, msg)
	if err != nil {
		return nil, err
	}

	cache.DisableCache()

	return res, nil
}

func GetMsg() (string, error) {
	if cache.Status == true {
		return cache.Data, nil
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	database := new(store.DatabaseMongo)
	connection := database.OpensConnection(ctx)
	defer database.CloseConnection(connection, ctx)

	res, err := database.GetMsg(connection, ctx)
	if err != nil {
		return "", err
	}

	cache.EnableCache(res)

	return res, nil
}
