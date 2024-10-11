package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func NewCrawlCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "crawl",
		Short: "crawl",
		Long:  "\n crawl data",
		Run: func(cmd *cobra.Command, args []string) {
			option := GetOptionByViper()

			runner, err := NewHuawei(option.ConfigFile)
			if err != nil {
				os.Exit(1)
			}
			runner.Run()
		},
	}

	return cmd
}
