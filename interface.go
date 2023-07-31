package mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gone-io/gone"
)

const IdGoneMqtt = "gone-mqtt"

func NewClient() (gone.Goner, gone.GonerId) {
	return &client{}, IdGoneMqtt
}

type client struct {
	gone.Flag
	mqtt.Client

	url      string `gone:"config,mqtt.url"`
	username string `gone:"config,mqtt.username"`
	password string `gone:"config,mqtt.password"`
}

func (c *client) Start(gone.Cemetery) error {
	opts := mqtt.
		NewClientOptions().
		AddBroker(c.url).
		SetUsername(c.username).
		SetPassword(c.password)

	c.Client = mqtt.NewClient(opts)

	if token := c.Client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return nil
}

func (c *client) Stop(gone.Cemetery) error {
	c.Client.Disconnect(2000)
	return nil
}
