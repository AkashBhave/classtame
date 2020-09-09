package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var prefix string

// workCmd represents the work command
var workCmd = &cobra.Command{
	Use:   "work",
	Short: "List a class's work in Taskwarrior",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Prefix set in viper overrides default of 'edu'
		if viper.IsSet("prefix") {
			prefix = viper.GetString("prefix")
		}

		// Run Taskwarrior
		command := exec.Command("task", fmt.Sprintf("pro:%s.%s", prefix, args[0]),
			"rc.verbose=label",
			"rc._forcecolor:on")
		stdout, err := command.Output()
		if err != nil {
			fmt.Println("There is no work for the specified class")
		}
		fmt.Print(string(stdout))
	},
}

func init() {
	rootCmd.AddCommand(workCmd)

	workCmd.Flags().StringVar(&prefix, "prefix", "edu", "Taskwarrior project prefix")
}
