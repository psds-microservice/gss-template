package cmd

import (
	"fmt"
	"log"

	"gss/internal/tools"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install or clone template repository",
	Long:  "Clone template repo into config directory. Set GSS_TEMPLATE_REPO or template_repository_url in config.",
	RunE:  runInstall,
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func runInstall(cmd *cobra.Command, args []string) error {
	url := viper.GetString("template_repository_url")
	if url == "" {
		url = viper.GetString("GSS_TEMPLATE_REPO")
	}
	if url == "" {
		log.Fatal("template_repository_url or GSS_TEMPLATE_REPO not set. Configure in ~/.config/gss/config.yaml or env.")
	}

	tplPath := viper.GetString("template_repository_path")
	if tplPath == "" {
		log.Fatal("template_repository_path not set in config.")
	}

	if err := tools.MkdirAll(tplPath); err != nil {
		return err
	}

	if tools.IsGitRepo(tplPath) {
		fmt.Println("Template already installed. Use 'gss update' to pull latest.")
		return nil
	}

	if err := tools.RunCommand("git", []string{"clone", url, "."}, tplPath); err != nil {
		return err
	}
	fmt.Println("Template installed.")
	return nil
}
