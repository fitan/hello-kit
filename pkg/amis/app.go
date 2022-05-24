package amis

type App struct {
	Pages []Children `json:"pages"`
}

type Children struct {
	Label     string     `json:"label"`
	URL       string     `json:"url,omitempty"`
	Children  []Children `json:"children,omitempty"`
	SchemaAPI string     `json:"schemaApi,omitempty"`
	Rewrite   string     `json:"rewrite,omitempty"`
	Icon      string     `json:"icon,omitempty"`
}

type Page struct {
	Type    string      `json:"type"`
	Title   string      `json:"title"`
	Remark  interface{} `json:"remark"`
	Name    string      `json:"name"`
	Toolbar []interface{}
	Body    []interface{} `json:"body"`
}

type CURDBody struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	Api    string `json:"api"`
	Filter struct {
		Title         string `json:"title"`
		Mode          string `json:"mode"`
		WrapWithPanel bool   `json:"wrapWithPanel"`
		SubmitText    string `json:"submitText"`
		Controls      []struct {
			Type        string `json:"type"`
			Name        string `json:"name"`
			Placeholder string `json:"placeholder"`
			AddOn       struct {
				Label     string `json:"label"`
				Type      string `json:"type"`
				ClassName string `json:"className"`
			} `json:"addOn"`
			Clearable bool `json:"clearable"`
		} `json:"controls"`
		ClassName string `json:"className"`
	} `json:"filter"`
	BulkActions []struct {
		Label      string `json:"label"`
		Type       string `json:"type"`
		ActionType string `json:"actionType"`
		Level      string `json:"level"`
		Dialog     struct {
			Title string `json:"title"`
			Name  string `json:"name"`
			Body  struct {
				Type     string `json:"type"`
				Api      string `json:"api"`
				Controls []struct {
					Type  string `json:"type"`
					Name  string `json:"name"`
					Label string `json:"label"`
				} `json:"controls"`
			} `json:"body"`
		} `json:"dialog,omitempty"`
		Api         string `json:"api,omitempty"`
		ConfirmText string `json:"confirmText,omitempty"`
	} `json:"bulkActions"`
	Columns []struct {
		Name     string      `json:"name,omitempty"`
		Label    string      `json:"label"`
		Sortable bool        `json:"sortable,omitempty"`
		Width    interface{} `json:"width,omitempty"`
		Type     string      `json:"type,omitempty"`
		Buttons  []struct {
			Type    string `json:"type"`
			Buttons []struct {
				Type        string `json:"type"`
				Label       string `json:"label"`
				Level       string `json:"level"`
				ActionType  string `json:"actionType"`
				Link        string `json:"link,omitempty"`
				ConfirmText string `json:"confirmText,omitempty"`
				Api         string `json:"api,omitempty"`
			} `json:"buttons"`
		} `json:"buttons,omitempty"`
		Placeholder string `json:"placeholder,omitempty"`
		Fixed       string `json:"fixed,omitempty"`
	} `json:"columns"`
	AffixHeader      bool   `json:"affixHeader"`
	ColumnsTogglable string `json:"columnsTogglable"`
	Placeholder      string `json:"placeholder"`
	TableClassName   string `json:"tableClassName"`
	HeaderClassName  string `json:"headerClassName"`
	FooterClassName  string `json:"footerClassName"`
	ToolbarClassName string `json:"toolbarClassName"`
	CombineNum       int    `json:"combineNum"`
	BodyClassName    string `json:"bodyClassName"`
}
