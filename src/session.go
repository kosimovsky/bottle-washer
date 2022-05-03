package src

import (
	"fmt"
	"github.com/stmcginnis/gofish"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/netip"
	"os"
)

type authConf struct {
	Ip      []netip.Addr `yaml:"ip,flow"`
	Net     string       `yaml:"network"`
	User    string       `yaml:"user"`
	Passwd  string       `yaml:"passwd"`
	InsFlag bool         `yaml:"insecure"`
}

func (ac *authConf) readAuthFile(cfg *string) {
	file, err := ioutil.ReadFile(*cfg)
	if err != nil {
		log.Printf("couldn't read file: %s", *cfg)
	}
	err = yaml.Unmarshal(file, ac)
	if err != nil {
		log.Fatalf("couldnt't unmarshal yaml data: %v", err)
	}
}

type Client struct {
	ApiClient *gofish.APIClient
}

func InitClient(config *string) (p *Client) {
	var t Client
	t.ApiClient = InitClientWConfig(config)
	return &t
}

func Config(cfg *string) (conf gofish.ClientConfig) {
	var ac authConf
	ac.readAuthFile(cfg)
	//fmt.Println(ac.Ip[0].String())
	conf = gofish.ClientConfig{
		Endpoint: "https://" + ac.Ip[0].String(),
		Username: ac.User,
		Password: ac.Passwd,
		Insecure: ac.InsFlag,
	}
	return conf
}

func InitClientWConfig(cfg *string) *gofish.APIClient {
	config := Config(cfg)
	c, err := gofish.Connect(config)
	if err != nil {
		err := fmt.Errorf("failure while connect to %v", config.Endpoint)
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return c
}
