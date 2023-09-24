package main

const CONF_PATH = "./conf/config.json"

type inner struct {
	address string `json:"broker"`
	port    uint16 `json:"port"`
}

type cloud struct {
	url      string `json:"address"`
	protocol string `json:"protocol"`
}

type conf struct {
	broker inner
	host   cloud
}

var conf_info conf
