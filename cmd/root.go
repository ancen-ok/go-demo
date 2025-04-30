package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-demo/router"
	"net/http"
)

var Param CommandParam

var defaultPort int = 8683

type CommandParam struct {
	Port     int64
	Config   string
	Username string
	Version  string
}

var rootCmd = &cobra.Command{
	Use:   "help",
	Short: "帮助你快速配置服务",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server start ")
		router.Run(Param.Port)
	},
}

func startServer(port int) {
	fmt.Printf("服务启动中，端口: %d", port)
	addr := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Printf("启动服务失败: %v", err)
	}
}

func init() {
	rootCmd.PersistentFlags().Int64VarP(&Param.Port, "port", "p", 8089, "服务端口")
	rootCmd.PersistentFlags().StringVarP(&Param.Config, "config", "c", "config.yaml", "配置文件")
	rootCmd.PersistentFlags().StringVarP(&Param.Version, "version", "v", "v1.0.0", "版本号")
	rootCmd.PersistentFlags().StringVarP(&Param.Username, "username", "u", "admin", "用户名")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
