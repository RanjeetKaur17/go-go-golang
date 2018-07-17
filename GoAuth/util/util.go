package util

import (
	"os/exec"
	"log"
	"strings"
)

func GetUUID() string {

	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(out))
}
