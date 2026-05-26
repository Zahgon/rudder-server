package alert

func (ops *VictorOps) Alert(message string) { _ = "STUB: not implemented"; return }

// Not handling errors when sending alert to victorops

type VictorOps struct {
	routingKey   string
	instanceName string
}
