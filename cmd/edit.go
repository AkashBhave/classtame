package cmd

import (
	"errors"
	"fmt"

	"akashbhave.com/classtame/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit information about a class",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Class doesn't exist
		if !viper.IsSet(fmt.Sprintf("classes.%s", args[0])) {
			return errors.New("Class doesn't exist, use 'new'")
		}

		var class util.Class
		viper.UnmarshalKey(fmt.Sprintf("classes.%s", args[0]), &class)

		// Update fields
		if period != "" {
			class.Period = period
		}
		if link != "" {
			class.Link = link
		}
		if name != "" {
			class.Name = name
		}
		if teacher != "" {
			class.Teacher = teacher
		}

		viper.Set(fmt.Sprintf("classes.%s", args[0]), class)

		// If write was successful, return nil
		err := viper.WriteConfig()
		if err == nil {
			fmt.Println("Edit successful")
		}
		return err
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Target variables were defined in new.go
	editCmd.Flags().StringVarP(&period, "period", "p", "", "period of the class")
	editCmd.Flags().StringVarP(&link, "link", "l", "", "link to the virtual classroom")
	editCmd.Flags().StringVar(&name, "name", "", "full name of the class")
	editCmd.Flags().StringVar(&teacher, "teacher", "", "the teacher's name")
}
