package cmd

import (
	"fmt"
	"log"
	"strings"

	"gss/internal/scaffold"
	"gss/internal/tools"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	CIGitHub = "github"
	CIGitLab = "gitlab"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create new project from template",
	Long:  "Scaffold new Go service project. Run 'gss install' first if template is not yet installed.",
	RunE:  runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("group", "g", "", "Git group/org (e.g. myorg)")
	initCmd.Flags().StringP("project", "p", "", "Project name (e.g. my-service)")
	initCmd.Flags().StringP("kuber", "k", "", "Kubernetes namespace")
	initCmd.Flags().String("ci", "", "CI/CD: github or gitlab")
}

func runInit(cmd *cobra.Command, args []string) error {
	group, _ := cmd.Flags().GetString("group")
	project, _ := cmd.Flags().GetString("project")
	kuber, _ := cmd.Flags().GetString("kuber")
	ci, _ := cmd.Flags().GetString("ci")

	if group == "" {
		if err := survey.AskOne(&survey.Input{Message: "Git group/org:", Default: "myorg"}, &group); err != nil {
			return err
		}
	}
	if project == "" {
		if err := survey.AskOne(&survey.Input{Message: "Project name:", Default: "my-service"}, &project); err != nil {
			return err
		}
	}
	if kuber == "" {
		if err := survey.AskOne(&survey.Input{Message: "Kubernetes namespace:", Default: "default"}, &kuber); err != nil {
			return err
		}
	}
	if ci == "" {
		if err := survey.AskOne(&survey.Select{
			Message: "CI/CD:",
			Options: []string{CIGitLab, CIGitHub},
			Default: CIGitLab,
		}, &ci); err != nil {
			return err
		}
	}
	ci = strings.ToLower(ci)
	if ci != CIGitHub && ci != CIGitLab {
		ci = CIGitLab
	}

	tplPath := viper.GetString("template_path")
	if tplPath == "" {
		log.Fatal("template_path not set. Run 'gss install' or check config.")
	}
	if !tools.Exists(tplPath) {
		log.Fatalf("template not found at %s. Run 'gss install' first.", tplPath)
	}

	s, err := scaffold.New(tplPath, group, project, kuber, ci)
	if err != nil {
		return err
	}

	if err := s.Run(); err != nil {
		return err
	}
	fmt.Println("Project scaffolded successfully.")
	return nil
}
