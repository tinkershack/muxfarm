package cmd

import (
	"github.com/spf13/cobra"
)

// senseCmd represents the sense command
var stitchCmd = &cobra.Command{
	Use:   "stitch",
	Short: "Holds a set of runners to stitch operations",
	Long:  `Holds a set of runners to stitch operations`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("stitch called")
	// },
}

func init() {
	rootCmd.AddCommand(stitchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stitchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stitchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
