package src

import (
	"fmt"
	"net"
	"net/netip"
	"os"
	"sort"
	"strconv"
	"time"
)

type Concurrent struct {
	ConcurrentMax int
}

func (c *Concurrent) SetMaxConcurrency() {
	if val, ok := os.LookupEnv("CONCURRENTMAX"); ok {
		c.ConcurrentMax, _ = strconv.Atoi(val)
	} else {
		c.ConcurrentMax = 255
	}
}

type alive struct {
	addr      netip.Addr
	available bool
}

func hostArray(cidr string) ([]netip.Addr, error) {
	prefix, err := netip.ParsePrefix(cidr)
	if err != nil {
		return nil, fmt.Errorf("error while parsing CIDR %v", err)
	}
	var ips []netip.Addr
	for addr := prefix.Addr(); prefix.Contains(addr); addr = addr.Next() {
		ips = append(ips, addr)
	}
	if len(ips) < 2 {
		return ips, nil
	}
	return ips[1 : len(ips)-1], err
}

func recieveAnsw(quantity int, resChan <-chan alive, readyChan chan<- []alive) {
	var alives []alive
	for i := 0; i < quantity; i++ {
		res := <-resChan
		if res.available {
			alives = append(alives, res)
		}
	}
	readyChan <- alives
}

func tryHost(reqChan <-chan netip.Addr, resChan chan<- alive) {
	for ip := range reqChan {
		ipPort := net.JoinHostPort(ip.String(), "443")
		var avlbl bool
		_, err := net.DialTimeout("tcp4", ipPort, 3*time.Second)
		if err != nil {
			avlbl = false
		} else {
			avlbl = true
		}
		//err = conn.Close()
		resChan <- alive{addr: ip, available: avlbl}
	}

}

func CheckAlivesInCidr(cidr string) ([]netip.Addr, error) {
	hosts, err := hostArray(cidr)
	if err != nil {
		return nil, err
	}
	var c Concurrent
	c.SetMaxConcurrency()
	concurrentMax := c.ConcurrentMax
	reqChan := make(chan netip.Addr, concurrentMax)
	resChan := make(chan alive, len(hosts))
	readyChan := make(chan []alive)
	for _, ip := range hosts {
		reqChan <- ip
	}
	for i := 0; i < concurrentMax; i++ {
		go tryHost(reqChan, resChan)
	}
	go recieveAnsw(len(hosts), resChan, readyChan)
	alives := <-readyChan
	return aliveToArray(alives), nil
}

func CheckAlivesInIPs(ar []netip.Addr) ([]netip.Addr, error) {
	var c Concurrent
	c.SetMaxConcurrency()
	concurrentMax := c.ConcurrentMax
	reqChan := make(chan netip.Addr, concurrentMax)
	resChan := make(chan alive, len(ar))
	readyChan := make(chan []alive)
	for _, ip := range ar {
		reqChan <- ip
	}
	for i := 0; i < concurrentMax; i++ {
		go tryHost(reqChan, resChan)
	}
	go recieveAnsw(len(ar), resChan, readyChan)
	alives := <-readyChan
	return aliveToArray(alives), nil
}

func aliveToArray(alives []alive) []netip.Addr {
	var ret []netip.Addr
	for _, a := range alives {
		ret = append(ret, a.addr)
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Less(ret[j])
	})
	return ret
}

func Check(c Client, args []string) {
	if len(args) > 1 {
		switch args[1] {
		case "cidr":
			ar, err := CheckAlivesInCidr(args[2])
			if err != nil {
				fmt.Println(err.Error())
			}
			printAddr(ar)
		}
	} else {
		err := fmt.Errorf("%s command need subcommand, args count %d", "connect", len(args))
		fmt.Println(err.Error())
	}
	defer c.ApiClient.Logout()
}

func printAddr(ar []netip.Addr) {
	for _, addr := range ar {
		fmt.Println(addr)
	}
}
