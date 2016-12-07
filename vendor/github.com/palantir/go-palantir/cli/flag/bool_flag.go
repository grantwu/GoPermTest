package flag

import (
	"strconv"
)

type BoolFlag struct {
	Name       string
	Alias      string
	Value      bool
	Usage      string
	Deprecated string
}

func (f BoolFlag) MainName() string {
	return f.Name
}

func (f BoolFlag) FullNames() []string {
	if f.Alias == "" {
		return []string{WithPrefix(f.Name)}
	}
	return []string{WithPrefix(f.Name), WithPrefix(f.Alias)}
}

func (f BoolFlag) IsRequired() bool {
	return false
}

func (f BoolFlag) DeprecationStr() string {
	return f.Deprecated
}

func (f BoolFlag) HasLeader() bool {
	return true
}

func (f BoolFlag) Default() interface{} {
	return f.Value
}

func (f BoolFlag) Parse(str string) (interface{}, error) {
	return strconv.ParseBool(str)
}

func (f BoolFlag) PlaceholderStr() string {
	panic("bool flag does not have placeholder")
}

func (f BoolFlag) DefaultStr() string {
	return ""
}

func (f BoolFlag) EnvVarStr() string {
	return ""
}

func (f BoolFlag) UsageStr() string {
	return f.Usage
}
