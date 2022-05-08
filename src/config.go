// Package src
/*
Copyright © 2022 Alexander Kosimovsky a.kosimovsky@gmail.com

*/
package src

import (
	"fmt"
	"github.com/stmcginnis/gofish"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/netip"
)

type AuthConf struct {
	Ip      []netip.Addr `yaml:"ip,flow"`
	Net     string       `yaml:"network"`
	User    string       `yaml:"user"`
	Passwd  string       `yaml:"passwd"`
	InsFlag bool         `yaml:"insecure"`
}

var CfgFile string

const Logfile = "bottle-washer.log"

func (ac *AuthConf) ReadAuthFile() {
	cfg := CfgFile
	file, err := ioutil.ReadFile(cfg)
	if err != nil {
		log.Printf("couldn't read file: %s", cfg)
	}
	err = yaml.Unmarshal(file, ac)
	if err != nil {
		log.Fatalf("couldnt't unmarshal yaml data: %v", err)
	}
}

func (ac *AuthConf) CountClients() (int, error) {
	if ac.Net != "" {
		ar, err := CheckAlivesInCidr(ac.Net)
		if err != nil {
			return 0, err
		}
		return len(ar), err
	}
	if len(ac.Ip) > 0 {
		ar, err := CheckAlivesInIPs(ac.Ip)
		if err != nil {
			return 0, err
		}
		return len(ar), err
	}
	return 0, nil
}

func (ac *AuthConf) ClientConfig() (ret []gofish.ClientConfig) {
	if ac.Net != "" {
		ar, err := CheckAlivesInCidr(ac.Net)
		if err != nil {
			fmt.Printf("error in ParseConf net %v\n", err.Error())
		}
		for _, addr := range ar {
			conf := gofish.ClientConfig{
				Endpoint: "https://" + addr.String(),
				Username: ac.User,
				Password: ac.Passwd,
				Insecure: ac.InsFlag,
			}
			ret = append(ret, conf)
		}
	} else if len(ac.Ip) > 0 {
		ar, err := CheckAlivesInIPs(ac.Ip)
		if err != nil {
			fmt.Printf("error in ParseConf ips %v\n", err.Error())
		}
		for _, addr := range ar {
			conf := gofish.ClientConfig{
				Endpoint: "https://" + addr.String(),
				Username: ac.User,
				Password: ac.Passwd,
				Insecure: ac.InsFlag,
			}
			ret = append(ret, conf)
		}
	}
	return ret
}

//func InitClients(clientChan chan<- Client, configs []gofish.ClientConfig) {
//	for _, config := range configs {
//		clientChan <- InitClientWConfig(config)
//	}
//}

type Client struct {
	ApiClient *gofish.APIClient
}

func InitClientWConfig(cfg gofish.ClientConfig) (client Client) {
	var err error
	client.ApiClient, err = gofish.Connect(cfg)
	if err != nil {
		err := fmt.Errorf("failure while connect to %v", cfg.Endpoint)
		log.Print(err.Error())
	}
	return client
}
