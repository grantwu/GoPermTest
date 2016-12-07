package cli

type Manpage struct {
	Source     string // e.g. Linux
	Manual     string // e.g. Linux Programmer's Manual
	BugTracker string
	SeeAlso    []ManpageRef
}

type ManpageRef struct {
	Name    string
	Section uint8
}
