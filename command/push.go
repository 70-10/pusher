package command

import (
	"fmt"
	"io/ioutil"

	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/certificate"
)

type PushCommand struct {
	Meta
}

func (c *PushCommand) Run(args []string) int {
	err := c.Config.InitializePayload()
	if err != nil {
		return ExitCodeInitializePayloadError
	}

	cert, err := certificate.FromP12File(c.Config.P12FilePath, c.Config.P12Password)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Certificate Error: %v", err))
		return ExitCodeCertificateError
	}

	notification := &apns2.Notification{}
	notification.DeviceToken = c.Config.DeviceToken
	notification.Topic = c.Config.Topic
	payload, err := ioutil.ReadFile(c.Config.PayloadFilePath)
	if err != nil {
		return ExitCodePayloadFileError
	}
	notification.Payload = payload

	client := apns2.NewClient(cert)
	env := c.Config.Env
	if env == "production" || env == "prod" || env == "pro" {
		client = client.Production()
	} else if env == "development" || env == "develop" || env == "dev" {
		client = client.Development()
	} else {
		c.Ui.Error("Environment is unexpected")
		return ExitCodeClientEnvironmentError
	}

	res, err := client.Push(notification)
	if err != nil {
		c.Ui.Error(fmt.Sprintf("Push Notification Error: %v", err))
		return ExitCodePushNotificationError
	}

	c.Ui.Output(fmt.Sprintf("%v %v %v", res.StatusCode, res.ApnsID, res.Reason))
	return ExitCodeOK
}

func (c *PushCommand) Synopsis() string {
	return "Push Notification"
}

func (c *PushCommand) Help() string {
	return ""
}
