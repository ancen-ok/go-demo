package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-demo/router"
)

var serverCmd = &cobra.Command{
	Use:               "server",
	Short:             "Start server",
	DisableAutoGenTag: false,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server start ")
		router.Run(Param.Port)
	},
}

func init() {
	rootCmd.Flags().Int64VarP(&Param.Port, "port", "p", 8282, "服务端口")
	rootCmd.Flags().StringVarP(&Param.Config, "config", "c", "config.yaml", "配置文件")
	rootCmd.AddCommand(serverCmd)
}
