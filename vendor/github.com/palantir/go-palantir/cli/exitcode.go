package cli

type ExitCoder interface {
	error
	ExitCode() int
}

type exitCodeError struct {
	error
	exitCode int
}

func (e *exitCodeError) ExitCode() int {
	return e.exitCode
}

func WithExitCode(exitCode int, err error) ExitCoder {
	return &exitCodeError{
		exitCode: exitCode,
		error:    err,
	}
}
