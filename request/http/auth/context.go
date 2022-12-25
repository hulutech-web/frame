package auth

type Context interface {
	AuthClaimID() (ID int64, exist bool)
	IUserModel() IUser
}
