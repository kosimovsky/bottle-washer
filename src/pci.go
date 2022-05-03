package src

import (
	"fmt"
)

func Pci(cfg *string, args []string) {
	newClient := InitClient(cfg)

	switch args[1] {
	case "all":
		newClient.listPCIDevices()
	case "endpoints":
		newClient.pciEndpoints()
	}
	defer newClient.ApiClient.Logout()
}

func (c *Client) listPCIDevices() {
	s, _ := c.ApiClient.Service.Systems()
	pciDevices, _ := s[0].PCIeDevices()
	for i, dev := range pciDevices {
		fmt.Printf("Device %d:\t%v\n", i+1, dev.Name)
	}
}

func (c *Client) pciEndpoints() {
	s, _ := c.ApiClient.Service.Systems()
	pciDevices, _ := s[0].PCIeDevices()
	for i, dev := range pciDevices {
		fmt.Printf("Endpoint %d:\t%v\n", i+1, dev.ODataID)
	}
}
