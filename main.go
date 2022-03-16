package main

import (
	"fmt"
	"gohub/app/cmd"
	"gohub/app/cmd/make"
	"gohub/bootstrap"
	btsConfig "gohub/config"
	"gohub/pkg/config"
	"gohub/pkg/console"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	btsConfig.Initialize()
}

func main() {

	var rootCmd = &cobra.Command{
		Use:   config.Get("app.name"),
		Short: "A Simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {
			// 配置初始化，依赖命令行 --env 参数
			config.InitConfig(cmd.Env)

			// 初始化 Logger
			bootstrap.SetupLogger()

			//初始化数据库
			bootstrap.SetupDB()

			//初始化redis
			bootstrap.SetupRedis()
		},
	}

	//注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
		make.CmdMake,
	)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	//执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}

}
