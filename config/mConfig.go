package config

type MConfig struct {
	Host           string   `yaml:"host"`
	Port           int      `yaml:"port"`
	Storage        mStorage `yaml:"storage"`
	Template       string   `yaml:"template"`
	UpdateEndpoint string   `yaml:"update_endpoint"`
	UpdateSecret   string   `yaml:"update_secret"`
	CachePath      string   `yaml:"cache_path"`

	MSite mSite `yaml:"site"`
}

var MConfigInstance *MConfig
