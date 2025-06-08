package model

type CardProperties struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Options []Options `json:"options"`
	Type    string    `json:"type"`
}
type Options struct {
	Color string `json:"color"`
	ID    string `json:"id"`
	Value string `json:"value"`
}
