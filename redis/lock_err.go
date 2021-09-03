package myredis

const (
	LockExistError   = "user lock alreadt exist"
	NotSameLockError = "unlock target is not expected"
)

func GetLockExistError() string {
	return LockExistError
}

func GetNotSameLockError() string {
	return NotSameLockError
}
