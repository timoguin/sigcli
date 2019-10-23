package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/signalsciences/go-sigsci"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listTopAttacksCmd = &cobra.Command{
	Use:   "list-top-attacks",
	Short: "List top attacks for a specific site",
	Run: func(cmd *cobra.Command, args []string) {
		email = viper.GetString("email")
		token = viper.GetString("token")
		corp = viper.GetString("corp")
		site = viper.GetString("site")
		query = viper.GetString("query")
		sc := sigsci.NewTokenClient(email, token)
		queryParsed, err := url.ParseQuery(query)
		if err != nil {
			log.Fatal(err)
		}
		ret, err := sc.ListTopAttacks(corp, site, queryParsed)
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
	rootCmd.AddCommand(listTopAttacksCmd)
	listTopAttacksCmd.Flags().StringVarP(&corp, "corp", "c", "", "Name of the corp")
	listTopAttacksCmd.Flags().StringVarP(&site, "site", "s", "", "Name of the site")
	listTopAttacksCmd.Flags().StringVarP(&query, "query", "q", "", "Query string to filter the results")
	viper.BindPFlag("corp", listTopAttacksCmd.Flags().Lookup("corp"))
	viper.BindPFlag("site", listTopAttacksCmd.Flags().Lookup("site"))
	viper.BindPFlag("query", listTopAttacksCmd.Flags().Lookup("query"))
}
