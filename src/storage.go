// Package src
/*
Copyright © 2022 Alexander Kosimovsky a.kosimovsky@gmail.com

*/
package src

import (
	"fmt"
	"math"
)

func Storage(c Client, args []string) {

	if len(args) > 1 {
		switch args[1] {
		case "ctrl":
			c.getControllers()
		case "pd":
			c.getPhysDisks()
		}
	} else {
		err := fmt.Errorf("%s command need subcommand, args count %d", "storage", len(args))
		fmt.Println(err.Error())
	}
}

func (c *Client) getControllers() {
	s, _ := c.ApiClient.Service.Systems()
	storage, _ := s[0].Storage()

	//for _, d := range drives {
	//	fmt.Println(d.Manufacturer, d.PartNumber, d.Model, d.SerialNumber, int(float64(d.CapacityBytes)/math.Pow(1000, 3)))
	//}

	for _, ctrl := range storage {
		fmt.Println(ctrl.Name, ctrl.Status.State, ctrl.DrivesCount)
	}
	defer c.ApiClient.Logout()
}

func (c *Client) getPhysDisks() {
	s, _ := c.ApiClient.Service.Systems()
	ctrl, _ := s[0].Storage()
	for _, s := range ctrl {
		if s.DrivesCount > 0 {
			fmt.Printf("Controller: %s\n\n", s.Name)
			c, _ := s.Drives()
			for i, d := range c {
				fmt.Printf("Disk%d:\n\tManufacturer: %s\n\tModel: %s\n\tSerial Number: %s\n\tCapacity: %dGB\n", i,
					d.Manufacturer, d.Model, d.SerialNumber, int(float64(d.CapacityBytes)/math.Pow(1000, 3)))
			}
		}
	}
	defer c.ApiClient.Logout()
}

//func listDisks(d redfish.Drive) {
//	println(d.Manufacturer, d.PartNumber, d.Model, d.SerialNumber, int(float64(d.CapacityBytes)/math.Pow(1000, 3)))
//}
