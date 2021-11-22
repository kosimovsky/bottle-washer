package proxmox_ssh_key

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

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

	scanner := bufio.NewScanner(f)
	line := 1
	for scanner.Scan() {
		b, _ := regexp.Match("(?<= name: )([a-z0-9\\-]+)", []byte(scanner.Text()))
		if  b {
			fmt.Println(b)
		}
		line++
	}
}