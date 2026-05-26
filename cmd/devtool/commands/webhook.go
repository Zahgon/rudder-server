package commands

import (
	"net/http"

	"github.com/urfave/cli/v2"
)

func init() {
	DefaultList = append(DefaultList, WEBHOOK())
}

func WEBHOOK() *cli.Command { _ = "STUB: not implemented"; return nil }

func WebhookRun(c *cli.Context) error { _ = "STUB: not implemented"; return nil }

type webhook struct {
	Verbose bool
}

type payload struct {
	SentAt string
}

func (*webhook) computeTime(b []byte) { _ = "STUB: not implemented"; return }

func (wh *webhook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
