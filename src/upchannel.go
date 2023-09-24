package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// ==============================================================================
var up_client mqtt.Client

// ==============================================================================
func onUpConnected(client mqtt.Client) {

}

func onUpDisconnected(client mqtt.Client, err error) {

}

func onUpMessage(client mqtt.Client, message mqtt.Message) {

}

func up_init() {
	var err error
	up_client, err = connectToBroker(onUpConnected, onUpDisconnected, "upchannel")
	if err != nil {
		fmt.Println("connect to inner broker fails: ", err)
	}
}

func transferTo(addr string, port uint16) {

	for {

	}
}
