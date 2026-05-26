package controlplane

import (
	"context"

	proto "github.com/rudderlabs/rudder-server/proto/common"
)

type AuthInfo struct {
	Service         string
	ConnectionToken string
	InstanceID      string
	TokenType       string
	Labels          map[string]string
}

type authService struct {
	authInfo AuthInfo
	proto.UnimplementedDPAuthServiceServer
}

func (a *authService) GetConnectionToken(_ context.Context, _ *proto.GetConnectionTokenRequest) (*proto.GetConnectionTokenResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *authService) GetWorkspaceToken(_ context.Context, _ *proto.GetWorkspaceTokenRequest) (*proto.GetWorkspaceTokenResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
