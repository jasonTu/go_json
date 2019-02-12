package main

import (
	"fmt"
	"regexp"
	"strings"
)

func CheckIpField(ip string) (status bool) {
	regIpv4 := regexp.MustCompile(`^(25[0-5]|2[0-4]\d|[0-1]?\d?\d)(\.(25[0-5]|2[0-4]\d|[0-1]?\d?\d)){3}$`)
	regIpv6 := regexp.MustCompile(`^(([\da-fA-F]{1,4}):){8}$`)
	ips := strings.Split(ip, ",")
	for _, sip := range ips {
	    sip = strings.TrimSpace(sip)
		if regIpv4.MatchString(sip) || regIpv6.MatchString(sip) {
			continue
		} else {
			fmt.Println("CheckIpField: ip field invalid ->", sip)
			return false
		}
	}
	return true
}

func main() {
	// str := "10.1.2.23, 10.1.2.23, 114.179.206.171"
	str := "10.1.2.23"
	/*
	ips := strings.Split(str, ",")
	for _, ip := range ips {
		fmt.Println(ip)
		fmt.Println(strings.TrimSpace(ip))
	}
	*/
	CheckIpField(str)
	str = "10.1.2.23, 10.1.2.23, 114.179.206.171"
	CheckIpField(str)
	str = "10.1.2.2311, 10.1.2.23, 114.179.206.171"
	CheckIpField(str)
}
