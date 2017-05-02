package command

const (
	ExitCodeOK = iota
	ExitCodeInitializeConfigError
	ExitCodeRunCommandError
	ExitCodeCertificateError
	ExitCodeClientEnvironmentError
	ExitCodePushNotificationError
	ExitCodeInitializePayloadError
	ExitCodePayloadFileError
)
