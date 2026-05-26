package commands

import (
	"github.com/urfave/cli/v2"
)

var DefaultList []*cli.Command

func init() {
	DefaultList = append(DefaultList, ETCD())
}

func ETCD() *cli.Command { _ = "STUB: not implemented"; return nil }

func Mode(c *cli.Context) error { _ = "STUB: not implemented"; return nil }

func List(c *cli.Context) error { _ = "STUB: not implemented"; return nil }
