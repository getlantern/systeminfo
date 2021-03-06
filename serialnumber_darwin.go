package systeminfo

import (
	"fmt"
	"os/exec"
	"regexp"
)

var (
	regex = regexp.MustCompile(`"IOPlatformSerialNumber"\s*=\s*"(.+)"`)
)

// SerialNumber returns the device's serial number
func SerialNumber() (string, error) {
	cmd := exec.Command("ioreg", "-l")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error running ioreg: %v", err)
	}
	groups := regex.FindStringSubmatch(string(out))
	if len(groups) != 2 {
		return "", fmt.Errorf("unexpected number of matcing groups from ioreg command output: %d", len(groups))
	}
	return groups[1], nil
}
