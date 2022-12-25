package saas

type databaser interface {
	ConnectionArgs() string
	Driver() string
	Prefix() string
}
