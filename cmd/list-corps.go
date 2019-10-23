package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/signalsciences/go-sigsci"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/*
var (
	email     string
    token     string
)
*/

var listCorpsCmd = &cobra.Command{
	Use:   "list-corps",
	Short: "List sigsci corps",
	Run: func(cmd *cobra.Command, args []string) {
		// Get dah corps
		email = viper.GetString("email")
		token = viper.GetString("token")
		sc := sigsci.NewTokenClient(email, token)
		corps, err := sc.ListCorps()
		if err != nil {
			log.Fatal(err)
		}
		output, err := json.MarshalIndent(corps, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", output)
	},
}

func init() {
	rootCmd.AddCommand(listCorpsCmd)
	listCorpsCmd.Flags().StringVarP(&email, "email", "e", "", "Sigsci username/email")
	listCorpsCmd.Flags().StringVarP(&token, "token", "t", "", "Sigsci API token")
	viper.BindPFlag("email", listCorpsCmd.Flags().Lookup("email"))
	viper.BindPFlag("token", listCorpsCmd.Flags().Lookup("token"))
}
