package alert

var (
	alertProvider       string
	pagerDutyRoutingKey string
	instanceName        string
	victorOpsRoutingKey string
)

func Init() { _ = "STUB: not implemented"; return }

func loadConfig() { _ = "STUB: not implemented"; return }

// AlertManager interface
type AlertManager interface {
	Alert(string)
}

// New returns FileManager backed by configured privider
func New() (AlertManager, error) { _ = "STUB: not implemented"; return *new(AlertManager), nil }
