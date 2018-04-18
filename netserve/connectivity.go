package netserve

import (
	"net"
	"time"
)

func IsTCPAlive(service string) bool {

	dialTO := 1000
	conn, err := net.DialTimeout("tcp", service, time.Duration(dialTO) * time.Millisecond)
	if err == nil {
		conn.Close()
		return true
	}

	return false
}
