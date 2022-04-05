package utils

import (
	"net"
)

func ConvertMacToInt() int64 {
	var sum int64 = 0
	macStr, err := getMac()
	if err == nil {
		for _, mac := range macStr {
			var macChar = mac
			sum += int64(macChar)
		}
	}
	return sum
}

func getMac() (string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	var macAddr = ""
	for _, netInterface := range netInterfaces {
		macAddr = netInterface.HardwareAddr.String()
		if len(macAddr) > 0 {
			break
		}
	}

	return macAddr, nil
}
