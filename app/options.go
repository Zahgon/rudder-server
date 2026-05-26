package app

// Options contains application's initialisation options
type Options struct {
	NormalMode      bool
	DegradedMode    bool
	ClearDB         bool
	Cpuprofile      string
	Memprofile      string
	VersionFlag     bool
	EnterpriseToken string
}

// LoadOptions loads application's initialisation options based on command line flags and environment
func LoadOptions(args []string) *Options { _ = "STUB: not implemented"; return nil }

// Parse command line options

// Ignore errors; flagSet is set for ExitOnError.
