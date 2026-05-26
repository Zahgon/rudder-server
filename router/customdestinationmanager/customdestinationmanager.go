package customdestinationmanager

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/sony/gobreaker"

	"github.com/rudderlabs/rudder-go-kit/logger"

	backendconfig "github.com/rudderlabs/rudder-server/backend-config"
)

const (
	STREAM = "stream"
	KV     = "kv"
)

var (
	ObjectStreamDestinations    []string
	KVStoreDestinations         []string
	Destinations                []string
	pkgLogger                   logger.Logger
	disableEgress               bool
	skipBackendConfigSubscriber bool
)

// DestinationManager implements the method to send the events to custom destinations
type DestinationManager interface {
	SendData(jsonData json.RawMessage, destID string) (int, string)
	BackendConfigInitialized() <-chan struct{}
}

// CustomManagerT handles this module
type CustomManagerT struct {
	destType    string
	managerType string

	stateMu  sync.RWMutex // protecting all 4 maps below
	config   map[string]backendconfig.DestinationT
	breaker  map[string]*breakerHolder
	clientMu map[string]*sync.RWMutex
	client   map[string]*clientHolder

	timeout                  time.Duration
	breakerTimeout           time.Duration
	backendConfigInitialized chan struct{}
}

// clientHolder keeps the config of a destination and corresponding producer for a stream destination
type clientHolder struct {
	config map[string]any
	client any
}

type breakerHolder struct {
	config    map[string]any
	breaker   *gobreaker.CircuitBreaker
	lastError error
}

func Init() { _ = "STUB: not implemented"; return }

func loadConfig() { _ = "STUB: not implemented"; return }

// newClient delegates the call to the appropriate manager
func (customManager *CustomManagerT) newClient(destID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (customManager *CustomManagerT) send(jsonData json.RawMessage, client any, config map[string]any) (int, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

// If client is not properly initialized then it won't reach here

// if the event supports HSET operation then use HSET

// SendData gets the producer from streamDestinationsMap and sends data
func (customManager *CustomManagerT) SendData(jsonData json.RawMessage, destID string) (int, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

func (customManager *CustomManagerT) close(destID string) { _ = "STUB: not implemented"; return }

func (customManager *CustomManagerT) onNewDestination(destination backendconfig.DestinationT) error {
	_ = "STUB: not implemented" // skipcq: CRT-P0003
	return nil
}

func (customManager *CustomManagerT) onConfigChange(destID string, newDestConfig map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

type Opts struct {
	Timeout time.Duration
}

// New returns CustomdestinationManager
func New(destType string, o Opts) DestinationManager {
	_ = "STUB: not implemented"
	return *new(DestinationManager)
}

func (customManager *CustomManagerT) BackendConfigInitialized() <-chan struct{} {
	_ = "STUB: not implemented"
	return nil
}

func (customManager *CustomManagerT) backendConfigSubscriber() { _ = "STUB: not implemented"; return }

func (customManager *CustomManagerT) genComparisonConfig(config any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}
