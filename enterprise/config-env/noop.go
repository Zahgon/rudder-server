package configenv

type NOOP struct{}

func (*NOOP) ReplaceConfigWithEnvVariables(workspaceConfig []byte) (updatedConfig []byte) {
	_ = "STUB: not implemented"
	return nil
}
