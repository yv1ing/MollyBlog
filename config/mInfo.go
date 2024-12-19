package config

type mInfo struct {
	Logo        string `yaml:"logo"`
	Link        string `yaml:"link"`
	Title       string `yaml:"title"`
	Author      string `yaml:"author"`
	Email       string `yaml:"email"`
	Language    string `yaml:"language"`
	Copyright   string `yaml:"copyright"`
	Description string `yaml:"description"`

	Motto []string `yaml:"motto"`
}
