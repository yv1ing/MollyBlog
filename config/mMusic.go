package config

type mMusic struct {
	Title string       `yaml:"title"`
	List  []mMusicItem `yaml:"list"`
}

type mMusicItem struct {
	Name     string `yaml:"name"`
	Singer   string `yaml:"singer"`
	Cover    string `yaml:"cover"`
	MusicUrl string `yaml:"music_url"`
	LyricUrl string `yaml:"lyric_url"`
}
