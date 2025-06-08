package model

type Timezone struct {
	AutomaticTimezone    string `json:"automaticTimezone"`
	ManualTimezone       string `json:"manualTimezone"`
	UseAutomaticTimezone bool   `json:"useAutomaticTimezone"`
}

type Props struct {
	CustomStatus string `json:"customStatus"`
}
type UserPatchProps struct {
	Props Props `json:"props"`
}

type CustomStatusProp struct {
	Emoji     string `json:"emoji"`
	Text      string `json:"text"`
	Duration  string `json:"duration"`
	ExpiresAt string `json:"expiresAt"`
}

func NewUserPatchProps() *UserPatchProps {
	return &UserPatchProps{}
}

type AccessToken struct {
	AccessToken string
}
