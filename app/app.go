package app

//go:generate mockgen -destination=../mocks/app/mock_app.go -package=mock_app github.com/rudderlabs/rudder-server/app App

import (
	"net/http"
	"os"

	"github.com/rudderlabs/rudder-go-kit/logger"

	"github.com/rudderlabs/rudder-server/jobsdb"
)

const (
	GATEWAY   = "GATEWAY"
	PROCESSOR = "PROCESSOR"
	EMBEDDED  = "EMBEDDED"
)

// App represents a rudder-server application
type App interface {
	Setup()              // Initializes application
	Stop()               // Stop application
	Options() *Options   // Get this application's options
	Features() *Features // Get this application's enterprise features
}

// app holds the main application's configuration and state
type app struct {
	log      logger.Logger
	options  *Options
	features *Features // Enterprise features, if available

	cpuprofileOutput *os.File
}

// Setup initializes application
func (a *app) Setup() {
	_ = "STUB: not implemented"
	// If cpuprofile flag is present, setup cpu profiling
	return
}

func (a *app) initCPUProfiling() { _ = "STUB: not implemented"; return }

func (a *app) initFeatures() { _ = "STUB: not implemented"; return }

// Options returns this application's options
func (a *app) Options() *Options {
	_ = "STUB: not implemented"

	// Features returns this application's enterprise features
	return nil
}

func (a *app) Features() *Features {
	_ = "STUB: not implemented"

	// Stop stops application
	return nil
}

func (a *app) Stop() { _ = "STUB: not implemented"; return }

// get up-to-date statistics

// New creates a new application instance
func New(options *Options) App { _ = "STUB: not implemented"; return *new(App) }

// LivenessHandler is the http handler for the Kubernetes liveness probe
func LivenessHandler(jobsDB jobsdb.JobsDB) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

func getHealthVal(jobsDB jobsdb.JobsDB) (bool, string) { _ = "STUB: not implemented"; return false, "" }
