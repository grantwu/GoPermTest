package main

import (
	"github.com/Playground/PermTester/permtester"
	"github.com/palantir/go-palantir/cli"
	"github.com/palantir/go-palantir/cli/flag"
	"os"
	"github.com/palantir/stacktrace"
)

func main() {
	app := createApp()
	os.Exit(app.Run(os.Args))
}

func createApp() *cli.App {
	app := cli.NewApp()
	app.Name = "PermTester"
	app.Version = "0.1"

	app.Flags = []flag.Flag{
		flag.StringFlag{
			Name:  "file",
			Alias: "f",
			Usage: "fileToTest",
		},
	}

	app.Subcommands = []cli.Command{
		{
			Name:   "canread",
			Action: canRead,
		},
		{
			Name:   "canwrite",
			Action: canWrite,
		},
	}

	return app
}

func canRead(cliCtx cli.Context) error {
	permTester := permtester.NewLinuxPermTester()
	if canRead, err := permTester.CanRead(cliCtx.String("file")); err != nil {
		return stacktrace.Propagate(err, "")
	} else if canRead {
		println("Can read file.")
	} else {
		println("Cannot read file.")
	}

	return nil
}

func canWrite(cliCtx cli.Context) error {
	permTester := permtester.NewLinuxPermTester()
	if canWrite, err := permTester.CanWrite(cliCtx.String("file")); err != nil {
		return stacktrace.Propagate(err, "")
	} else if canWrite {
		println("Can write file.")
	} else {
		println("Cannot write file.")
	}

	return nil
}
