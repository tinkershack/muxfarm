package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tinkershack/muxfarm/stitch"
)

// ingestBatcherCmd represents the ingestBatcher command
var ingestSplitterCmd = &cobra.Command{
	Use:   "ingestSplitter",
	Short: "Split media inputs form IngestDoc",
	Long:  `Split media inputs form IngestDoc`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ingestSplitter called")
		stitch.IngestSplitter(context.Background(), args)
	},
}

func init() {
	stitchCmd.AddCommand(ingestSplitterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ingestSplitterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ingestSplitterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
