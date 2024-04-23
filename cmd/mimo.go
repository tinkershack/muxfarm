package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/tinkershack/muxfarm/server"
)

// mimoCmd represents the mimo command
var mimoCmd = &cobra.Command{
	Use:   "mimo",
	Short: "mimo serves media operations from client",
	Long:  `mimo serves media operations from client.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("mimo called")
		server.MIMO(context.Background(), args)
	},
}

func init() {
	rootCmd.AddCommand(mimoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mimoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mimoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
