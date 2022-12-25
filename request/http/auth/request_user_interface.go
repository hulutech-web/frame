package auth

type RequestIUser interface {
	ScanUser() error
	User() IUser
	UserId() (userI int64, err error)
	ScanUserWithJSON() (isAbort bool)
}
