/*
Package admin :
- has a rpc over http server listening on a unix socket
- support other packages to expose any admin functionality over the above server

# Example for registering admin handler from another package

// Add this while initializing the package in setup or init etc.
admin.RegisterAdminHandler("PackageName", &PackageAdmin{})
admin.RegisterStatusHandler("PackageName", &PackageAdmin{})

// Convention is to keep the following code in admin.go in the respective package
type PackageAdmin struct {
}

// Status function is used for debug purposes by the admin interface

	func (p *PackageAdmin) Status() map[string]interface{} {
		return map[string]interface{}{
			"parameter-1"  : value,
			"parameter-2"  : value,
		}
	}

// The following function can be called from rudder-cli using getUDSClient().Call("PackageName.SomeAdminFunction", &arg, &reply)

	func (p *PackageAdmin) SomeAdminFunction(arg *string, reply *string) error {
		*reply = "admin function output"
		return nil
	}
*/
package admin

import (
	"context"
	"net/rpc"

	"github.com/rudderlabs/rudder-go-kit/logger"
)

// RegisterAdminHandler is used by other packages to
// expose admin functions over the unix socket based rpc interface
func RegisterAdminHandler(name string, handler any) { _ = "STUB: not implemented"; return }

// @TODO fix ignored error

type Admin struct {
	rpcServer *rpc.Server
}

var (
	instance  *Admin
	pkgLogger logger.Logger
)

func Init() { _ = "STUB: not implemented"; return }

type LogLevel struct {
	Module string
	Level  string
}

func (*Admin) SetLogLevel(l LogLevel, reply *string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

//nolint:forbidigo

// GetLoggingConfig returns the logging configuration
func (*Admin) GetLoggingConfig(_ struct{}, reply *string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

//nolint:forbidigo

// StartServer starts an HTTP server listening on unix socket and serving rpc communication
func StartServer(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
