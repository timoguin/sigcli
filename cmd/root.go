package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cli config options, flags, params
var (
	cfgFile string
	debug   bool
	email   string
	token   string
	corp    string
	site    string
	query   string
)

// compile time build info
var (
	Version   string
	GitCommit string
	BuildTime string
	GoVersion string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sigcli",
	Short: "Interact with the Signal Sciences API",
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
	rootCmd.PersistentFlags().StringVar(&email, "email", "", "SigSci email/username")
	rootCmd.PersistentFlags().StringVar(&token, "token", "", "SigSci API token")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sigcli.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enable debug mode for verbose output")
	viper.BindPFlag("email", rootCmd.PersistentFlags().Lookup("email"))
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search for config
		viper.AddConfigPath("/etc/sigsci/")
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName(".sigsci")
	}

	viper.SetEnvPrefix("sigsci")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Printf("Found config file %s\n", viper.ConfigFileUsed())
	}
}
