package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

const BROKER_PATH string = "./"
const BROKER_NAME string = "mosquitto"

func checkBroker() {
	for {
		output, _ := exec.Command("ps", "-ef", BROKER_NAME).Output()
		fields := strings.Fields(string(output))

		living := false
		for _, v := range fields {
			if v == BROKER_NAME {
				living = true
				break
			}
		}

		time.Sleep(time.Second * 30)
		if false == living {
			fmt.Println("[warn]mosquitto exit, restart now")
		}
	}
}

func startBroker(args []string) bool {
	full_path := fmt.Sprintf("%s%s", BROKER_PATH, BROKER_NAME)
	err := syscall.Exec(full_path, args, os.Environ())
	if err != nil {
		return true
	} else {
		return false
	}
}

func stopBroker() {
	cmd := exec.Cmd{}
}
