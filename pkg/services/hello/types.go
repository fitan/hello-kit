package hello

type SayReq struct {
	Uri struct {
		ID int
	}
	Body SayRes
}

type SayRes struct {
	ID    int
	Name  string
	Email string
}
