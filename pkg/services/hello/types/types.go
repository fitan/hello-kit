package types

type SayReq struct {
	Uri struct {
		ID int `json:"id" uri:"id"`
	}
}

type SayRes struct {
	ID    int
	Name  string
	Email string
}
