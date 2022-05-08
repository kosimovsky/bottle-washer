/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bottle-washer/src"
	"github.com/spf13/cobra"
	"sync"
)

// storageCmd represents the storage command
var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Command to get storage entities as controllers, disks etc.",
	Long: `Usage: 
storage get ctrl	-- to get controllers
storage get pd		-- to get physical disks`,
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
				src.Storage(c, args)
			}(client)
		}
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(storageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// storageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	storageCmd.Flags().String("get", "", "Get storage collection")
}
