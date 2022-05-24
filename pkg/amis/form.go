package amis

type AmisForm []FormBody

type FormBody struct {
	Type string     `json:"type"`
	Api  string     `json:"api"`
	Mode string     `json:"mode"`
	Body []ItemForm `json:"body"`
}

type ItemForm struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Label       string `json:"label"`
	Required    bool   `json:"required"`
	Placeholder string `json:"placeholder"`
	Size        string `json:"size"`
}
