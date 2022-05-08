// Package src
/*
Copyright © 2022 Alexander Kosimovsky a.kosimovsky@gmail.com

*/
package src

func ClearJobs(c Client, args []string) {
	c.getJobs()
}

func (c *Client) getJobs() {
	s, err := c.ApiClient.Service.Managers()

	if err != nil {
		panic(err)
	}
	nn, _ := s[0].EthernetInterfaces()
	println(nn)
	defer c.ApiClient.Logout()
	//addrArr := nn[0].IPv4Addresses
	//for _, sys := range addrArr {
	//	println(sys.Address)
	//}
}
