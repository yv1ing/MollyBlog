package config

type mFriend struct {
	Title string        `json:"title"`
	List  []mFriendItem `json:"list"`
}

type mFriendItem struct {
	Name        string `yaml:"name"`
	Link        string `yaml:"link"`
	Avatar      string `yaml:"avatar"`
	Description string `yaml:"description"`
}
