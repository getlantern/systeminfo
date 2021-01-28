package systeminfo

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bi-zone/wmi"
)

var (
	regex = regexp.MustCompile(`"IOPlatformSerialNumber" = "(.+)"`)
)

type Win32_BIOS struct {
	SerialNumber string
}

// SerialNumber returns the device's serial number
func SerialNumber() (string, error) {
	var results []Win32_BIOS
	q := wmi.CreateQuery(&results, "")
	err := wmi.Query(q, &results)
	if err != nil {
		return "", err
	}
	if len(results) != 1 {
		return "", fmt.Errorf("unexpected number of BIOS WMI query results: %d", len(results))
	}
	serialNumber := strings.TrimSpace(results[0].SerialNumber)
	if serialNumber == "" {
		return "", fmt.Errorf("serial number is blank")
	}
	return serialNumber, nil
}
