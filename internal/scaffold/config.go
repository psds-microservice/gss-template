package scaffold

// Config is the scaffold config (scaffold.json / hero.json-like).
type Config struct {
	Command map[string][][]string `json:"command"`
	Copy    map[string]string     `json:"copy"`
	Gitkeep []string              `json:"gitkeep"`
	Mkdir   []string              `json:"mkdir"`
	Render  []RenderRule          `json:"render"`
}

// RenderRule describes one template → output file.
type RenderRule struct {
	From string `json:"from"`
	To   string `json:"to"`
	// Optional: only render when CI matches (e.g. "github", "gitlab"). Empty = always.
	CI string `json:"ci,omitempty"`
}
