package cmd

import (
	"github.com/quankori/go-hd-wallet/pkg/accounts"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var checkUrlCmd = &cobra.Command{
	Use:   "check-url",
	Short: "check-url",
	Long:  `check-url`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			log.Fatal("subcommand check-url only take one argument as url")
		}
		accounts.GenKey()
	},
}
