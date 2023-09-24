package main

import (
	"database/sql"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/mattn/go-sqlite3"
)

// ==============================================================================
var db_connection *sql.DB
var db_client mqtt.Client

// ==============================================================================
const CONFIG_DB_PATH = "./db/config.db3"
const DATA_DB_PATH = "./db/data.db3"

// ==============================================================================
func onDbConnected(client mqtt.Client) {
	var err error
	db_connection, err = sql.Open("sqlite3", CONFIG_DB_PATH)
	if err != nil {
		panic(err)
	}
}

func onDbDisconnected(client mqtt.Client, err error) {

}

func onDbMessage(client mqtt.Client, message mqtt.Message) {

}

func dbmgr_init() {
	var err error
	db_client, err = connectToBroker(onDbConnected, onDbDisconnected, "dbmgr")
	if err != nil {
		fmt.Println("[critical]dbmgr connect to broker fails: ", err)
	} else {
		fmt.Println("[info]dbmgr connect to broker success")
	}
}

func doQueryVersion(topic string) {
	// {
	// 	"token": "123",
	// 	"timestamp": "2019-03-01T09:30:08.230+0800",
	// 	"body": []
	// }
	payload := "{\"token\":\""
	payload += genToken() + "\",\"timestamp\":\""
	payload += genTimestamp() + "\",\"body\":[]}"
	db_client.Publish(topic, 0, true, payload)
}

func onQueryVersion(client mqtt.Client, message mqtt.Message) {
}

func doQueryModel(topic string) {
}

func onQueryModel(client mqtt.Client, message mqtt.Message) {
}

func doAddModel(topic string) {
}

func onAddModel(client mqtt.Client, message mqtt.Message) {
}

// var dbop = [...]mqtt_op{
// 	{"+/do/query/db/version", doQueryVersion, "db/on/response/+/version", onQueryVersion},
// 	{"+/do/query/db/model", doQueryModel, "db/on/response/+/model", onQueryModel},
// 	{"+/do/add/db/model", doAddModel, "db/response/+/model", onQueryModel},
// }

var dbop = make(map(string)mqtt_op, 10)

func dbmgr_run() {

	//1 subscribe
	for _, item := range dbop {
		db_client.Subscribe(item.responseTopic, 0, nil)
	}

	//2 loop
}
