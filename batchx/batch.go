package batchx

import (
	"context"
	"github.com/tao-kit/gormc"
	"gorm.io/gorm"
)

type BatchExecModel[DBModel any] interface {
	GetCacheKeys(data *DBModel) []string
	ExecCtx(ctx context.Context, execCtx gormc.ExecCtxFn, keys ...string) error
}

func BatchExecCtx[DBModel any, Model BatchExecModel[DBModel]](ctx context.Context, model Model, olds []DBModel, exec func(db *gorm.DB) error) error {
	if len(olds) == 0 {
		return nil
	}
	cacheKeys := getCacheKeysByMultiData(model, olds)

	err := model.ExecCtx(ctx, func(conn *gorm.DB) error {
		return exec(conn)
	}, cacheKeys...)
	return err
}

func getCacheKeysByMultiData[DBModel any, Model BatchExecModel[DBModel]](m Model, data []DBModel) []string {
	if len(data) == 0 {
		return []string{}
	}
	var keys []string
	for _, v := range data {
		keys = append(keys, m.GetCacheKeys(&v)...)
	}
	keys = uniqueKeys(keys)
	return keys
}
func uniqueKeys(keys []string) []string {
	keySet := make(map[string]struct{})
	for _, key := range keys {
		keySet[key] = struct{}{}
	}

	uniKeys := make([]string, 0, len(keySet))
	for key := range keySet {
		uniKeys = append(uniKeys, key)
	}

	return uniKeys
}
