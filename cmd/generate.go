package cmd

import (
	"fmt"
	"strconv"

	"github.com/quankori/go-hd-wallet/pkg/accounts"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var checkStatusCmd = &cobra.Command{
	Use:   "generate-wallet",
	Short: "generate-wallet for mnemonic key",
	Long:  `generate-wallet for mnemonic key`,
	Run: func(cmd *cobra.Command, args []string) {
		bs := []byte(strconv.Itoa(256))
		result, err := accounts.NewMnemonic(bs)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	},
}

func checkStatus(args []string) error {
	fmt.Printf("HI!! From check-status sub-command with %s as argument", args[0])
	return nil
}
