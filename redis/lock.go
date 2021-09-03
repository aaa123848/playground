package myredis

import (
	"context"
	"time"

	"github.com/google/uuid"
	dogError "github.com/iloveanimal/cutedog/errors"
)

func (redisStore *RedisStore) OnLock(ctx context.Context, target string) (string, error) {
	lockId := uuid.New().String()
	timeOut := 10 * time.Second

	status := redisStore.Db.SetNX(ctx, target, lockId, timeOut)
	if status.Err() != nil {
		return "", dogError.GetGeneralError(status.Err().Error())
	}
	if !status.Val() {
		return "", dogError.GetGeneralError(GetLockExistError())
	}
	return lockId, nil
}

func (redisStore *RedisStore) UnLock(ctx context.Context, target, lockId string) error {
	realLockId := redisStore.Db.Get(ctx, target).Val()
	if realLockId != lockId {
		return dogError.GetGeneralError(GetNotSameLockError())
	}

	res := redisStore.Db.Del(ctx, target)
	if res.Err() != nil {
		return dogError.GetGeneralError(res.Err().Error())
	}

	return nil
}

func (redisStore *RedisStore) CheckLockAvailable(ctx context.Context, target, lockId string) (bool, error) {
	realLockId := redisStore.Db.Get(ctx, target).Val()
	if realLockId != lockId {
		return false, nil
	}
	return true, nil
}
