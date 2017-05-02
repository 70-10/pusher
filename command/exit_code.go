package command

const (
	ExitCodeOK = iota
	ExitCodeCreateConfigError
	ExitCodeRunCommandError
	ExitCodeCertificateError
	ExitCodeClientEnvironmentError
	ExitCodePushNotificationError
	ExitCodePayloadFileError
)
