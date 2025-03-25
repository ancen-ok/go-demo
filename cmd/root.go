package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type CommandParam struct {
}

var (
	Port     int64
	Config   string
	Username string
	Version  string
	Home     string
)

var rootCmd = &cobra.Command{
	Use:   "help",
	Short: "帮助你快速配置服务",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("应用程序启动，端口: %d, 用户名: %s, verbose: %v\n", Port, Config, Version)
	},
}

func init() {
	rootCmd.PersistentFlags().Int64VarP(&Port, "port", "p", 8080, "服务端口")
	rootCmd.PersistentFlags().StringVarP(&Config, "config", "c", "config.yaml", "配置文件")
	rootCmd.PersistentFlags().StringVarP(&Version, "version", "v", "v1.0.0", "版本号")
	rootCmd.PersistentFlags().StringVarP(&Home, "home", "h", "home", "目录")
	rootCmd.PersistentFlags().StringVarP(&Username, "username", "u", "admin", "用户名")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
