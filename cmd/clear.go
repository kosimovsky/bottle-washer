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

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
				src.ClearJobs(c, args)
			}(client)
		}
		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
