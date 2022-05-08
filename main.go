/*
Copyright © 2022 Alexander Kosimovsky a.kosimovsky@gmail.com

*/
package main

import (
	"bottle-washer/cmd"
	"bottle-washer/src"
	"log"
	"os"
)

func main() {
	//Set logging
	logfile, err := os.OpenFile(src.Logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error while opening file : %v, %s", err, src.Logfile)
	}
	defer logfile.Close()
	log.New(logfile, "", log.LstdFlags|log.Lshortfile)
	log.SetOutput(logfile)
	//logging

	cmd.Execute()
}
