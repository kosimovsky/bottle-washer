package src

import (
	"fmt"
)

func Pci(c Client, args []string) {
	switch args[1] {
	case "all":
		c.listPCIDevices()
	case "endpoints":
		c.pciEndpoints()
	}
}

func (c *Client) listPCIDevices() {
	s, _ := c.ApiClient.Service.Systems()
	pciDevices, _ := s[0].PCIeDevices()
	for i, dev := range pciDevices {
		fmt.Printf("Device %d:\t%v\n", i+1, dev.Name)
	}
	defer c.ApiClient.Logout()
}

func (c *Client) pciEndpoints() {
	s, _ := c.ApiClient.Service.Systems()
	pciDevices, _ := s[0].PCIeDevices()
	for i, dev := range pciDevices {
		fmt.Printf("Endpoint %d:\t%v\n", i+1, dev.ODataID)
	}
	defer c.ApiClient.Logout()
}
