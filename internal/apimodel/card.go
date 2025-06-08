package apimodel

// Card 代表一个卡片的API模型
type Card struct {
	ID         string                 `json:"id"`
	Title      string                 `json:"title"`
	BoardID    string                 `json:"boardId"`
	Fields     []string               `json:"fields,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
}

type CardMoved struct {
	ID         string            `json:"id"`
	BoardID    string            `json:"boardId"`
	Title      string            `json:"title"`
	Properties map[string]string `json:"properties"`
}

type CardWithAslead struct {
	ID        string `json:"id"`
	BoardID   string `json:"boardId"`
	Title     string `json:"title"`
	AsleadId  string `json:"asleadId"`
	Status    string `json:"status"`
	OldStatus string `json:"oldStatus"`
}

type CardGroupLists struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Members []string `json:"members"`
}

type ListCards struct {
	Count       int
	StatusId    string
	StatusTitle string
	Members     []string
}
