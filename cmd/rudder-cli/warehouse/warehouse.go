package warehouse

import (
	"github.com/urfave/cli/v2"
)

type QueryResult struct {
	Columns []string
	Values  [][]string
}

type QueryInput struct {
	DestID       string
	SourceID     string
	SQLStatement string
}

type ConfigurationTestInput struct {
	DestID string
}

type ConfigurationTestOutput struct {
	Valid bool
	Error string
}

func Query(c *cli.Context) (err error) { _ = "STUB: not implemented"; return nil }

func ConfigurationTest(c *cli.Context) (err error) { _ = "STUB: not implemented"; return nil }
