package permtester

import (
	"os"
	"syscall"
	"github.com/palantir/stacktrace"
)

type LinuxPermTester struct {

}

func NewLinuxPermTester() PermTester {
	return &LinuxPermTester{}
}

func (p *LinuxPermTester) CanRead(file string) (bool, error) {
	stat, err := os.Stat(file)
	if err != nil {
		return false, stacktrace.Propagate(err, "Failed to get fileInfo from: %s", file)
	}
	fm := stat.Mode()
	if fm & (1 << 2) != 0 {
		return true, nil
	} else if (fm & (1 << 5)) != 0 && (os.Getegid() == int(stat.Sys().(*syscall.Stat_t).Gid)) {
		return true, nil
	} else if (fm & (1 << 8)) != 0 && (os.Geteuid() == int(stat.Sys().(*syscall.Stat_t).Uid)) {
		return true, nil
	}
	return false, nil
}

func (p *LinuxPermTester) CanWrite(file string) (bool, error) {
	return false, nil
}