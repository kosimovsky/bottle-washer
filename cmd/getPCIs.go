// Package cmd
/*
Copyright © 2022 Alexander Kosimovsky a.kosimovsky@gmail.com

*/
package cmd

import (
	"bottle-washer/src"
	"github.com/spf13/cobra"
	"sync"
)

// getPCIsCmd represents the getPCIs command
var getPCIsCmd = &cobra.Command{
	Use:   "pci",
	Short: "Command to get pci entities",
	Long: `Usage:
pci list all		-- to get what in pci slots
pci list endpoints	-- to get PCI RedFish API endpoints`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var ac src.AuthConf
		ac.ReadAuthFile()
		configs := ac.ClientConfig()
		clients := make(chan src.Client)
		var wg sync.WaitGroup
		go func() {
			for _, conf := range configs {
				clients <- src.InitClientWConfig(conf)
			}
			close(clients)
		}()

		for client := range clients {
			wg.Add(1)
			go func(c src.Client) {
				defer wg.Done()
				src.Pci(c, args)
			}(client)
		}
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(getPCIsCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getPCIsCmd.PersistentFlags().String("foo", "", "A help for foo")
	getPCIsCmd.Flags().String("list", "", `List information about pci devices:
Examples:
	pci list all ---- lists endpoints and names for all pci devices`)
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//getPCIsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
