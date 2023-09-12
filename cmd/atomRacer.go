package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tinkershack/muxfarm/stitch"
)

// atomRacerCmd represents the atomRacer command
var atomRacerCmd = &cobra.Command{
	Use:   "atomRacer",
	Short: "Race the atom to the restaurant at the end of the universe where time stops",
	Long: `It's a single player race. If it dies, it starts over; no battons passed.

	Race the atom to the restaurant at the end of the universe where time stops`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("atomRacer called")
		stitch.AtomRacer(context.Background(), args)
	},
}

func init() {
	stitchCmd.AddCommand(atomRacerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// atomRacerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// atomRacerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
