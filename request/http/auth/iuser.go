package auth

type IUser interface {
	Scan(userId int64) error
	Value() interface{}
}
