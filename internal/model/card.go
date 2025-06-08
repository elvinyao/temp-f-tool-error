package model

import "github.com/mattermost/focalboard/server/model"

type CardPatchParamerters struct {
	CustomCardPath *model.CardPatch
}

func NewCardPatchParamerters() *CardPatchParamerters {
	customCardPatch := &model.CardPatch{}
	return &CardPatchParamerters{
		CustomCardPath: customCardPatch,
	}
}

func (cpp *CardPatchParamerters) Copy(cardPatch *model.Card) {
	cpp.CustomCardPath.UpdatedProperties = make(map[string]any)
	stringmaps := cardPatch.Properties

	for k, v := range stringmaps {
		cpp.CustomCardPath.UpdatedProperties[k] = v
	}
}

func (cpp *CardPatchParamerters) UpdateValue(key string, newValue any) {
	cpp.CustomCardPath.UpdatedProperties[key] = newValue
}

type CardRevised struct {
	ID                   string            `json:"id"`
	BoardID              string            `json:"boardId"`
	Title                string            `json:"title"`
	Properties           map[string]string `json:"properties"`
	IdToNameProps        map[string]string `json:"idToNameProps"`
	OptionIdToValueProps map[string]string `json:"optionIdToValueProps"`
}

type CardWithSpecifiedProperty struct {
	ID         string            `json:"id"`
	BoardID    string            `json:"boardId"`
	Title      string            `json:"title"`
	Properties map[string]string `json:"properties"`
}

type GroupCategoryIds []GroupCategoryId

type GroupCategoryId struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func InitGroupCategoryId() *GroupCategoryId {
	return &GroupCategoryId{}
}

type CardFromGroupName struct {
	ID          string
	Title       string
	Description string
	ListTitle   string
}

type BoardWithSpecifiedCardProperty struct {
	ID             string
	TeamID         string
	ChannelID      string
	Title          string
	Description    string
	CardProperties []map[string]interface{}
}
