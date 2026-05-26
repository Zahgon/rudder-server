package commands

import (
	"embed"

	"github.com/urfave/cli/v2"
)

func init() {
	DefaultList = append(DefaultList, EVENT())
}

//go:embed payloads/*
var payloads embed.FS

func EVENT() *cli.Command { _ = "STUB: not implemented"; return nil }

func EventSend(c *cli.Context) error { _ = "STUB: not implemented"; return nil }
