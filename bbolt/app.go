package bbolt

import (
	"github.com/urfave/cli/v2"
)

var (
	pathFlag = cli.StringFlag{
		Name:    "path",
		Aliases: []string{"p"},
		Value:   "",
		Usage:   "bbolt db file path",
	}
	bucketFlag = cli.StringFlag{
		Name:    "bucket",
		Aliases: []string{"b"},
		Value:   "",
		Usage:   "bbolt bucket name",
	}
	keyFlag = cli.StringFlag{
		Name:    "key",
		Aliases: []string{"k"},
		Value:   "",
		Usage:   "bbolt key",
	}
)

func NewApp(version string) *cli.App {
	app := cli.NewApp()

	app.Name = "bbolt-reader"
	app.Usage = "get bbolt value by specified key"
	app.Version = version
	app.Commands = []*cli.Command{
		GetCommand(),
	}

	return app
}

func GetCommand() *cli.Command {
	return &cli.Command{
		Name:      "get",
		Aliases:   []string{"g", "read"},
		ArgsUsage: "key",
		Usage:     "bbolt-reader get --path=<bbolt db path> --bucket=<bucket> --key=<key>",
		Action:    Get,
		Flags: []cli.Flag{
			&pathFlag,
			&bucketFlag,
			&keyFlag,
		},
	}
}
