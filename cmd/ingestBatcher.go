package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tinkershack/muxfarm/stitch"
)

// ingestBatcherCmd represents the ingestBatcher command
var ingestBatcherCmd = &cobra.Command{
	Use:   "ingestBatcher",
	Short: "Batch media inputs form IngestDoc",
	Long:  `Batch media inputs form IngestDoc`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ingestBatcher called")
		stitch.IngestBatcher(context.Background(), args)
	},
}

func init() {
	stitchCmd.AddCommand(ingestBatcherCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ingestBatcherCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ingestBatcherCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
