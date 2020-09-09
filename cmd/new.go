package cmd

import (
	"errors"
	"fmt"

	"akashbhave.com/classtame/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	period  int
	link    string
	name    string
	teacher string
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new class",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Class has already been created
		if viper.IsSet(fmt.Sprintf("classes.%s", args[0])) {
			return errors.New("Specified class already exists, use 'edit'")
		}

		// Create the class and add it to 'classes'
		class := &util.Class{
			Name:    name,
			Period:  period,
			Link:    link,
			Teacher: teacher,
		}
		viper.Set(fmt.Sprintf("classes.%s", args[0]), class)

		// If write was successful, return nil
		err := viper.WriteConfig()
		return err
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Required flags
	newCmd.Flags().IntVarP(&period, "period", "p", 0, "period of the class")
	newCmd.MarkFlagRequired("period")
	newCmd.Flags().StringVarP(&link, "link", "l", "", "link to the virtual classroom")
	newCmd.MarkFlagRequired("link")

	// Optional flags
	newCmd.Flags().StringVar(&name, "name", "", "full name of the class")
	newCmd.Flags().StringVar(&teacher, "teacher", "", "the teacher's name")
}
