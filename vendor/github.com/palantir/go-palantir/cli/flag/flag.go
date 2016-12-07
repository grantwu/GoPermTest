package flag

type Flag interface {
	MainName() string    // without hyphens
	FullNames() []string // with hyphens
	IsRequired() bool
	DeprecationStr() string
	HasLeader() bool

	Default() interface{} // must not be called if flag is required
	Parse(string) (interface{}, error)

	PlaceholderStr() string // must not be called on bool flag
	DefaultStr() string     // must not be called if flag is required
	EnvVarStr() string      // must not be called if flag is required
	UsageStr() string
}
