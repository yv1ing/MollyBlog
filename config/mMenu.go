package config

type mMenu struct {
	Items []mItem `yaml:"items"`
}

type mItem struct {
	Name string `yaml:"name"`
	Icon string `yaml:"icon"`
	Url  string `yaml:"url"`
}
