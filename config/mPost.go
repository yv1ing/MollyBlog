package config

type mPost struct {
	TocTitle   string      `yaml:"toc_title"`
	RecentPost mRecentPost `yaml:"recent_post"`
	Archive    mArchive    `yaml:"archive"`
	Tag        mTag        `yaml:"tag"`
}

type mRecentPost struct {
	Title  string `yaml:"title"`
	Number int    `yaml:"number"`
}

type mArchive struct {
	Title  string `yaml:"title"`
	Number int    `yaml:"number"`
}

type mTag struct {
	Number int `yaml:"number"`
}
