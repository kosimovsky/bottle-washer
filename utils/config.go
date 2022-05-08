package utils

import (
	"bottle-washer/cmd"
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

func (ac *AuthConf) ReadAuthFile() {
	cfg := cmd.CfgFile
	file, err := ioutil.ReadFile(cfg)
	if err != nil {
		log.Printf("couldn't read file: %s", cfg)
	}
	err = yaml.Unmarshal(file, ac)
	if err != nil {
		log.Fatalf("couldnt't unmarshal yaml data: %v", err)
	}
}

func Config(cfgChan chan<- []gofish.ClientConfig) {
	var ac AuthConf
	ac.ReadAuthFile()
	//conf = gofish.ClientConfig{
	//	Endpoint: "https://" + ac.Ip[0].String(),
	//	Username: ac.User,
	//	Password: ac.Passwd,
	//	Insecure: ac.InsFlag,
	//}
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

func ParseConfig() {

}
