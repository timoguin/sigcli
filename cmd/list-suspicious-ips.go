package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/signalsciences/go-sigsci"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listSuspiciousIPsCmd = &cobra.Command{
	Use:   "list-suspicious-ips",
	Short: "List suspicious IPS for a specific site",
	Run: func(cmd *cobra.Command, args []string) {
		// Get dah agents
		email = viper.GetString("email")
		token = viper.GetString("token")
		corp = viper.GetString("corp")
		site = viper.GetString("site")
		sc := sigsci.NewTokenClient(email, token)
		ret, err := sc.ListSuspiciousIPs(corp, site)
		if err != nil {
			log.Fatal(err)
		}
		output, err := json.MarshalIndent(ret, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", output)
	},
}

func init() {
	rootCmd.AddCommand(listSuspiciousIPsCmd)
	listSuspiciousIPsCmd.Flags().StringVarP(&corp, "corp", "c", "", "Name of the corp")
	listSuspiciousIPsCmd.Flags().StringVarP(&site, "site", "s", "", "Name of the site")
	viper.BindPFlag("corp", listSuspiciousIPsCmd.Flags().Lookup("corp"))
	viper.BindPFlag("site", listSuspiciousIPsCmd.Flags().Lookup("site"))
}
