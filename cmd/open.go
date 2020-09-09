package cmd

import (
	"errors"
	"fmt"

	"akashbhave.com/classtame/util"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a class's classroom link in the browser",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Class doesn't exist
		if !viper.IsSet(fmt.Sprintf("classes.%s", args[0])) {
			return errors.New("Class doesn't exist, use 'new'")
		}

		var class util.Class
		viper.UnmarshalKey(fmt.Sprintf("classes.%s", args[0]), &class)
		if class.Link == "" {
			return errors.New("Link not set on class, use 'edit'")
		}
		err := browser.OpenURL(class.Link)

		return err
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
