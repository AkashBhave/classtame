package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "classtame",
	Short: "A simple CLI utility to manage online classes",
	Long: `Classtame is a simple CLI utility to manage online classes.
The primary unit of this utility is a class.
You can perform actions related to a class with:
	classtame [command] [class]
Start by creating a new class with:
	classtame new [class]`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.classtame.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Search config in home directory with name ".classtame" (without extension)
		viper.AddConfigPath(home)
		viper.SetConfigName(".classtame")
		viper.SetConfigType("yaml")
		// Save the empty config if it doesn't exist
		viper.SafeWriteConfig()
	}

	viper.AutomaticEnv() // Read in environment variables that match

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
