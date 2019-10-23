package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/signalsciences/go-sigsci"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listAgentsCmd = &cobra.Command{
	Use:   "list-agents",
	Short: "List sigsci agents",
	Run: func(cmd *cobra.Command, args []string) {
		// Get dah agents
		email = viper.GetString("email")
		token = viper.GetString("token")
		corp = viper.GetString("corp")
		site = viper.GetString("site")
		sc := sigsci.NewTokenClient(email, token)
		agents, err := sc.ListAgents(corp, site)
		if err != nil {
			log.Fatal(err)
		}
		output, err := json.MarshalIndent(agents, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", output)
	},
}

func init() {
	rootCmd.AddCommand(listAgentsCmd)
	listAgentsCmd.Flags().StringVarP(&corp, "corp", "c", "", "Name of the corp")
	listAgentsCmd.Flags().StringVarP(&site, "site", "s", "", "Name of the site")
	viper.BindPFlag("corp", listAgentsCmd.Flags().Lookup("corp"))
	viper.BindPFlag("site", listAgentsCmd.Flags().Lookup("site"))
}
