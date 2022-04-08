package schema

type NodeAnnotation struct {
	Create bool
	Update bool
	Read bool
	Delete bool
}

func (n NodeAnnotation) Name() string {
	return "rest"
}

type FieldAnnotation struct {
	Create bool
	Update bool
	Read bool
	QueryOp struct{
		EQ bool
		NEQ bool
		In bool
		NotIn bool
		GT bool
		GTE bool
		LT bool
		LTE bool
		Contains bool
		HasPrefix bool
		HasSuffix bool
		EqualFold bool
		ContainsFold bool
	}
}

func (f FieldAnnotation) Name() string {
	return "rest"
}

type QueryStrOp struct {
	EQ bool
	NEQ bool
	In bool
	NotIn bool
	GT bool
	GTE bool
	LT bool
	LTE bool
	Contains bool
	HasPrefix bool
	HasSuffix bool
	EqualFold bool
	ContainsFold bool	
}

type QueryBoolOp struct {
	EQ bool
	NEQ bool	
}

type QueryNumOp struct {
	EQ bool
	NEQ bool
	In bool
	NotIn bool
	GT bool
	GTE bool
	LT bool
	LTE bool	
}

type QueryTimeOp struct {
	EQ bool
	NEQ bool
	In bool
	NotIn bool
	GT bool
	GTE bool
	LT bool
	LTE bool
}

type QeruyOptionalOp struct {
	In bool
	NotIn bool
}
