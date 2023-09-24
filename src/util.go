package main

import (
	"fmt"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type mqtt_op struct {
	queryTopic string
	onQuery    func(string)
	period     time.Duration
	onResponse func(mqtt.Client, mqtt.Message)
}

func connectToBroker(onconnected func(mqtt.Client), ondisconnected func(mqtt.Client, error), session string) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", conf_info.broker.address, conf_info.broker.port))
	opts.SetClientID(session)
	opts.OnConnect = onconnected
	opts.OnConnectionLost = ondisconnected
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client, nil
}

func genToken() string {
	return strconv.Itoa(time.Now().Minute()) + strconv.Itoa(time.Now().Second())
}

func genTimestamp() string {
	return time.Now().String()
}
