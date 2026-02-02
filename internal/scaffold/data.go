package scaffold

import "github.com/iancoleman/strcase"

// TemplateData is passed into all templates (paths and file content).
type TemplateData struct {
	Group   string
	Project ProjectNames
	Kuber   string
	CI      string // "github" | "gitlab"
}

// ProjectNames holds project name in different cases.
type ProjectNames struct {
	C string // CamelCase
	K string // kebab-case
	S string // snake_case
}

// NewTemplateData builds TemplateData from group, project name, kuber and CI.
func NewTemplateData(group, projectName, kuber, ci string) TemplateData {
	return TemplateData{
		Group: group,
		Project: ProjectNames{
			C: strcase.ToCamel(projectName),
			K: strcase.ToKebab(projectName),
			S: strcase.ToSnake(projectName),
		},
		Kuber: kuber,
		CI:    ci,
	}
}
