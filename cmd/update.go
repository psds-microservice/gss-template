package cmd

import (
	"fmt"
	"log"

	"gss/internal/tools"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update template from repository",
	Long:  "Run git pull in template directory.",
	RunE:  runUpdate,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func runUpdate(cmd *cobra.Command, args []string) error {
	tplPath := viper.GetString("template_repository_path")
	if tplPath == "" {
		log.Fatal("template_repository_path not set in config.")
	}
	if !tools.IsGitRepo(tplPath) {
		log.Fatal("Template directory is not a git repo. Run 'gss install' first.")
	}
	if err := tools.RunCommand("git", []string{"pull", "origin", "master"}, tplPath); err != nil {
		return err
	}
	fmt.Println("Template updated.")
	return nil
}
