package main

import (
	"fmt"
	"net"
	"time"
)

func main(){
	for i:=1; i<=1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.DialTimeout("tcp", address, time.Second)
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}