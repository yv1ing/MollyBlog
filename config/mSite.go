package config

type mSite struct {
	Info       mInfo       `yaml:"info"`
	Menu       mMenu       `yaml:"menu"`
	Post       mPost       `yaml:"post"`
	About      mAbout      `yaml:"about"`
	Friend     mFriend     `yaml:"friend"`
	Statistics mStatistics `yaml:"statistics"`
}
