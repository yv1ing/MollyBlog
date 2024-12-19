package config

type mSite struct {
	Info  mInfo  `yaml:"info"`
	Menu  mMenu  `yaml:"menu"`
	Post  mPost  `yaml:"post"`
	About mAbout `yaml:"about"`
}
