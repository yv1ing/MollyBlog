package config

type mStorage struct {
	SRC  string    `yaml:"src"`
	DST  string    `yaml:"dst"`
	Type string    `yaml:"type"`
	COS  cosConfig `yaml:"cos"`
}

type cosConfig struct {
	Appid     string `yaml:"appid"`
	Bucket    string `yaml:"bucket"`
	Region    string `yaml:"region"`
	SecretId  string `yaml:"secret_id"`
	SecretKey string `yaml:"secret_key"`
	SavePath  string `yaml:"save_path"`
}
