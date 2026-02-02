package cmd

import (
	"os"
	"path/filepath"

	"gss/internal/tools"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "gss [commands] [v2]",
	Short: "Go service scaffold generator",
	Long:  "Generate base project structure from template. Add 'v2' to use v2 template.",
	Args:  cobra.MaximumNArgs(1),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return initConfig(cmd, args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("config", "", "config file (default $HOME/.config/gss/config.yaml)")
}

func initConfig(cmd *cobra.Command, args []string) error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	gssConfigDir := filepath.Join(configDir, "gss")
	tplPath := filepath.Join(gssConfigDir, "tpl")
	templatePath := tplPath
	if len(args) > 0 && args[0] == "v2" {
		templatePath = filepath.Join(tplPath, "v2")
	}

	viper.AddConfigPath(gssConfigDir)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
		_ = tools.MkdirAll(gssConfigDir)
		viper.Set("template_repository_url", "")
		viper.Set("template_repository_path", filepath.Join(gssConfigDir, "tpl"))
		viper.Set("template_path", templatePath)
		configPath := filepath.Join(gssConfigDir, "config.yaml")
		if err := viper.SafeWriteConfigAs(configPath); err != nil {
			return err
		}
	} else if len(args) > 0 && args[0] == "v2" {
		viper.Set("template_path", templatePath)
	}

	viper.AutomaticEnv()
	return nil
}
