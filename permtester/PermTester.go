package permtester

type PermTester interface {
	CanRead(file string) (bool, error)
	CanWrite(file string) (bool, error)
}




