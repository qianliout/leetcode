package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	ConfigFile = "config-file"
)

type Option struct {
	ConfigFile string
}

func GetOptionByViper() Option {
	return Option{
		ConfigFile: viper.GetString(ConfigFile),
	}
}

func AddOption(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().StringP(ConfigFile, "c", "", "config file")

	rootCmd.PersistentFlags().VisitAll(func(flag *pflag.Flag) {
		err := viper.BindPFlag(flag.Name, flag)
		if err != nil {
			panic(err)
		}
	})

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()
}
