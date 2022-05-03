/*
Copyright © 2022 Alexander Kosimovsky a.kosimovsky@gmail.com

*/
package main

import (
	"bottle-washer/cmd"
	"bottle-washer/utils"
)

func main() {
	var ac utils.AuthConf
	ac.ReadAuthFile()
	//countServToConnectTo := utils.ParseConfig()
	//countClients, err := ac.CountClients()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//authConfChan := make(chan []gofish.ClientConfig, countClients)

	cmd.Execute()
}
