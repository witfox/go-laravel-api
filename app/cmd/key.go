package cmd

import (
	"gohub/helpers"
	"gohub/pkg/console"

	"github.com/spf13/cobra"
)

//生成 app key
var CmdKey = &cobra.Command{
	Use:   "key",
	Short: "Generate App Key, will print the generated Key",
	Run:   runKeyGenerate,
	Args:  cobra.NoArgs,
}

func runKeyGenerate(cmd *cobra.Command, args []string) {
	console.Success("----------------")
	console.Success("App Key: ")
	console.Success(helpers.RandomString(32))
	console.Success("----------------")
	console.Warning("please go to .env file to change the APP_KEY option")
}
