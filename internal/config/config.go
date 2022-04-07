package config

type AppConfig struct {
	InProduction bool
	UseCache     bool
	// TemplateCache map[string]*template.Template
	// Session       *scs.SessionManager
}
