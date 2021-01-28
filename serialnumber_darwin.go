package systeminfo

import (
	"fmt"
	"os/exec"
	"regexp"
)

var (
	regex = regexp.MustCompile(`"IOPlatformSerialNumber" = "(.+)"`)
)

// SerialNumber returns the device's serial number
func SerialNumber() (string, error) {
	cmd := exec.Command("ioreg", "-l")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	groups := regex.FindStringSubmatch(string(out))
	if len(groups) != 2 {
		return "", fmt.Errorf("unexpected number of groups from ioreg command: %d", len(groups))
	}
	return groups[1], nil
}
