package net_practice

import (
	"fmt"
	"net"
	"strings"
)

func Ipv6ToHostName(ipv6 string) string {
	sb := strings.Builder{}
	instanceIp := net.ParseIP(ipv6)
	for i := 2; i < net.IPv6len; i += 2 {
		fmt.Printf("%x\n", instanceIp[i])
		fmt.Printf("%x\n", instanceIp[i+1])
		if i == 4 {
			sb.WriteString("p")
		}
		if instanceIp[i] != 0 {
			sb.WriteString(fmt.Sprintf("%x", instanceIp[i]))
			sb.WriteString(fmt.Sprintf("%.02x", instanceIp[i+1]))
		} else {
			sb.WriteString(fmt.Sprintf("%x", instanceIp[i+1]))
		}
		sb.WriteString("-")
	}
	fmt.Println(sb.String()[:sb.Len()-1])
	return sb.String()[:sb.Len()-1]
}

//func main() {
//	ip := "2605:340:cdb1:100:7a17:0364:172f:55dd"
//	Ipv6ToHostName(ip)
//}
