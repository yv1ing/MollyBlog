package config

type mPost struct {
	TocTitle   string      `yaml:"toc_title"`
	RecentPost mRecentPost `yaml:"recent_post"`
	Archive    mArchive    `yaml:"archive"`
	Search     mSearch     `yaml:"search"`
	Tag        mTag        `yaml:"tag"`
	Category   mCategory   `yaml:"category"`
}

type mRecentPost struct {
	Title  string `yaml:"title"`
	Number int    `yaml:"number"`
}

type mArchive struct {
	Title  string `yaml:"title"`
	Number int    `yaml:"number"`
}

type mSearch struct {
	Title  string `yaml:"title"`
	Number int    `yaml:"number"`
}

type mTag struct {
	Title  string `yaml:"title"`
	Number int    `yaml:"number"`
}

type mCategory struct {
	Title  string `yaml:"title"`
	Number int    `yaml:"number"`
}
