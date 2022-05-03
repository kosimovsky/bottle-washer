package src

func ClearJobs(cfg *string, args []string) {
	newClient := InitClient(cfg)
	newClient.getJobs()
	defer newClient.ApiClient.Logout()
}

func (c *Client) getJobs() {
	s, err := c.ApiClient.Service.Managers()

	if err != nil {
		panic(err)
	}
	nn, _ := s[0].EthernetInterfaces()
	println(nn)
	//addrArr := nn[0].IPv4Addresses
	//for _, sys := range addrArr {
	//	println(sys.Address)
	//}
}
