package alert

import (
	"github.com/rudderlabs/rudder-go-kit/logger"
)

var (
	pagerDutyEndPoint = "https://events.pagerduty.com/v2/enqueue"
	pkgLogger         logger.Logger
)

func (ops *PagerDuty) Alert(message string) { _ = "STUB: not implemented"; return }

// Not handling errors when sending alert to victorops

type PagerDuty struct {
	instanceName string
	routingKey   string
}
