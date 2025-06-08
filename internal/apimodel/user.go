package apimodel

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	FullName string `json:"full_name"`
	Props    Props  `json:"props"`
}

type Props struct {
	CustomStatus string `json:"customStatus"`
}

type StatusEmojiConfig struct {
	Status   string
	Emoji    string
	Text     string
	Duration string
}
