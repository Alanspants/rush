package main

import (
	"fmt"
	"strconv"
	"strings"
)

func validIPAddress(queryIP string) string {
	IPType := ""

	splitIPv4 := strings.Split(queryIP, ".")
	splitIPv6 := strings.Split(queryIP, ":")

	if len(splitIPv6) == 8 {
		IPType = "IPv6"
	} else if len(splitIPv4) == 4 {
		IPType = "IPv4"
	} else {
		return "Neither"
	}

	if IPType == "IPv4" {
		for _, IP := range splitIPv4 {
			if len(IP) > 1 && IP[0] == '0' {
				return "Neither"
			}
			IPnum, err := strconv.Atoi(IP)
			if err != nil || IPnum > 255 || IPnum < 0 {
				return "Neither"
			}
		}
		return "IPv4"
		return "IPv4"
	} else if IPType == "IPv6" {
		for _, IP := range splitIPv6 {
			if len(IP) > 4 {
				return "Neither"
			}
			for i := 0; i < len(IP); i++ {
				if !((IP[0] >= '0' && IP[0] <= '9') || (IP[0] >= 'a' && IP[0] <= 'f') || (IP[0] >= 'A' && IP[0] <= 'F')) {
					return "Neither"
				}
			}
		}
		return "IPv6"
	}
	return "Neither"
}

func main() {
	fmt.Println(validIPAddress("02001:0db8:85a3:0000:0000:8a2e:0370:7334"))
}
