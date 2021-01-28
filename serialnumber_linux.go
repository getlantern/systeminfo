package systeminfo

import (
	"fmt"
	"os/exec"
	"regexp"
)

var (
	regex = regexp.MustCompile(`Serial Number:\s*(.+)\s`)
)

// SerialNumber returns the device's serial number
func SerialNumber() (string, error) {
	cmd := exec.Command("dmidecode", "-t", "system")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error running dmidecode: %v", err)
	}
	groups := regex.FindStringSubmatch(string(out))
	if len(groups) != 2 {
		return "", fmt.Errorf("unexpected number of matcing groups from dmidecode command output: %d", len(groups))
	}
	return groups[1], nil
}
