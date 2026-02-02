package scaffold

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gss/internal/render"
	"gss/internal/tools"
)

const configFilename = "scaffold.json"

// Scaffold runs mkdir → gitkeep → copy → render → command.
type Scaffold struct {
	tplPath string
	data    TemplateData
	cfg     Config
}

// New loads config from tplPath and builds Scaffold.
func New(tplPath, group, projectName, kuber, ci string) (*Scaffold, error) {
	cfgPath := filepath.Join(tplPath, configFilename)
	b, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("read config %s: %w", cfgPath, err)
	}
	var cfg Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}
	return &Scaffold{
		tplPath: tplPath,
		data:    NewTemplateData(group, projectName, kuber, ci),
		cfg:     cfg,
	}, nil
}

// Run executes all steps in order.
func (s *Scaffold) Run() error {
	root := s.data.Project.K
	if err := s.mkdir(root); err != nil {
		return err
	}
	if err := s.gitkeep(root); err != nil {
		return err
	}
	if err := s.copy(root); err != nil {
		return err
	}
	if err := s.render(root); err != nil {
		return err
	}
	if err := s.command(root); err != nil {
		return err
	}
	return nil
}

func (s *Scaffold) mkdir(root string) error {
	for _, p := range s.cfg.Mkdir {
		dest := filepath.Join(root, render.PathRender(p, s.data))
		if err := tools.MkdirAll(dest); err != nil {
			return err
		}
	}
	return nil
}

func (s *Scaffold) gitkeep(root string) error {
	for _, p := range s.cfg.Gitkeep {
		dest := filepath.Join(root, render.PathRender(p, s.data))
		if err := tools.Gitkeep(dest); err != nil {
			return err
		}
	}
	return nil
}

func (s *Scaffold) copy(root string) error {
	for from, to := range s.cfg.Copy {
		src := filepath.Join(s.tplPath, from)
		dest := filepath.Join(root, render.PathRender(to, s.data))
		if err := tools.Copy(src, dest); err != nil {
			return err
		}
	}
	return nil
}

func (s *Scaffold) render(root string) error {
	for _, r := range s.cfg.Render {
		if r.CI != "" && r.CI != s.data.CI {
			continue
		}
		src := filepath.Join(s.tplPath, r.From)
		toPath := render.PathRender(r.To, s.data)
		dest := filepath.Join(root, toPath)
		if err := render.File(src, dest, s.data); err != nil {
			return err
		}
	}
	return nil
}

func (s *Scaffold) command(root string) error {
	for cmd, argsList := range s.cfg.Command {
		for _, args := range argsList {
			if err := tools.RunCommand(cmd, args, root); err != nil {
				return err
			}
		}
	}
	return nil
}
