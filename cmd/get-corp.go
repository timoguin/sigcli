package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/signalsciences/go-sigsci"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//var corp string
/*
var (
	email     string
    token     string
)
*/

var getCorpCmd = &cobra.Command{
	Use:   "get-corp",
	Short: "Get info about a corp",
	Run: func(cmd *cobra.Command, args []string) {
		// Get dah corps
		email = viper.GetString("email")
		token = viper.GetString("token")
		corp = viper.GetString("corp")
		sc := sigsci.NewTokenClient(email, token)
		corp, err := sc.GetCorp(corp)
		if err != nil {
			log.Fatal(err)
		}
		output, err := json.MarshalIndent(corp, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", output)
	},
}

func init() {
	rootCmd.AddCommand(getCorpCmd)
	getCorpCmd.Flags().StringVarP(&email, "email", "e", "", "Sigsci username/email")
	getCorpCmd.Flags().StringVarP(&token, "token", "t", "", "Sigsci API token")
	getCorpCmd.Flags().StringVarP(&corp, "corp", "c", "", "Name of the corp")
	viper.BindPFlag("email", getCorpCmd.Flags().Lookup("email"))
	viper.BindPFlag("token", getCorpCmd.Flags().Lookup("token"))
	viper.BindPFlag("corp", getCorpCmd.Flags().Lookup("corp"))
}
