package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use: "core-api",
	Short: "this api for news portal",
	Run: func(cmd *cobra.Command, args[] string){
		cmd.Run(startCmd, nil)
	},
}

func Execute(){
	cobra.CheckErr(rootCmd.Execute())
}

func init(){
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config File (default is .env)")
	rootCmd.AddCommand(startCmd)
	rootCmd.Execute()
}

func initConfig(){
	if cfgFile != ""{
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigFile(`.env`)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}