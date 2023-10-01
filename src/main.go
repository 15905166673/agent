package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// ==============================================================================

// ==============================================================================

func init() {
	if false == startBroker([]string{""}) {
		fmt.Println("start mosquitto fails")
		os.Exit(-1)
	}

	content, err := os.ReadFile(CONF_PATH)
	err = json.Unmarshal(content, &conf_info)
	if err != nil {
		fmt.Println("[error]unmarshal conf fails: ", err)
		return
	}

	dbmgr_init()
	up_init()
}

func main() {
	fmt.Println("proxy start...")
	for {

	}
}
