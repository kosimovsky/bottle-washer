package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
)

type Node struct {
	hostname	[]string
	ip			[]net.IP
}


func main()  {
	path := "./corosync.conf"

	f, err := os.OpenFile(path, os.O_RDONLY, 0640)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	var n Node

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		sc := []byte(scanner.Text())
		b, _ := regexp.Match("(?: name: )([a-z0-9-]+)", sc)
		addr, _ := regexp.Match("(?: ring0_addr: )([0-9.]+)", sc)
		if  b {
			n.hostname = append(n.hostname, string(scanner.Bytes()[10:]))
		}
		if addr {
			n.ip = append(n.ip, net.ParseIP(string(scanner.Bytes())[16:]))
		}
	}
	j := 1
	for i := 0; i < len(n.hostname); i++ {
		fmt.Printf("node %d -- %s, and its ip address is %v\n", j, n.hostname[i], n.ip[i])
		j++
	}
}