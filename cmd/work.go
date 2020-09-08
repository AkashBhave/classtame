package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// workCmd represents the work command
var workCmd = &cobra.Command{
	Use:   "work",
	Short: "List a class's work in Taskwarrior",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("work called")
	},
}

func init() {
	rootCmd.AddCommand(workCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
