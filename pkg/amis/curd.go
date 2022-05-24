package amis

type CURD struct {
	Type string `json:"type"`
	Body struct {
		Type         string `json:"type"`
		Api          string `json:"api"`
		SyncLocation bool   `json:"syncLocation"`
		Columns      []struct {
			Name  string `json:"name"`
			Label string `json:"label"`
		} `json:"columns"`
	} `json:"body"`
}

type CURDHook struct {
	CURDSlice []CURD
}

//func (C CURDHook) Generate(graph *gen.Graph) error {
//	for _,n := range graph.Nodes {
//		curd := CURD{}
//		curd.Type =
//	}
//}
