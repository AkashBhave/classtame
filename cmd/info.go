package cmd

import (
	"errors"
	"fmt"

	"akashbhave.com/classtame/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Return information about a class", Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Class doesn't exist
		if !viper.IsSet(fmt.Sprintf("classes.%s", args[0])) {
			return errors.New("Class doesn't exist, use 'new'")
		}

		var class util.Class
		viper.UnmarshalKey(fmt.Sprintf("classes.%s", args[0]), &class)

		if class.Name != "" {
			fmt.Printf("Name: %s\n", class.Name)
		}
		fmt.Printf("Alias: %s\n", args[0])
		if class.Period != 0 {
			fmt.Printf("Period: %d\n", class.Period)
		}
		if class.Link != "" {
			fmt.Printf("Link: %s\n", class.Link)
		}
		if class.Teacher != "" {
			fmt.Printf("Teacher: %s\n", class.Teacher)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
