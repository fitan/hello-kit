package schema

type NodeAnnotation struct {
	Create bool `json:"create,omitempty"`
	CreateMany bool `json:"createMany,omitempty"`
	Update bool `json:"update,omitempty"`
	UpdateMany bool `json:"updateMany,omitempty"`
	Read bool `json:"read,omitempty"`
	ReadMany bool `json:"readMany,omitempty"`
	Delete bool `json:"delete,omitempty"`
	DeleteMany bool `json:"deleteMany,omitempty"`

	PagingMust bool `json:"pagingMust,omitempty"`
	PagingLimitMax int64 `json:"pagingLimitMax,omitempty"`

	Order bool `json:"order,omitempty"`
}

func (n NodeAnnotation) Name() string {
	return "rest"
}

type EdgeAnnotation struct {
	Create bool `json:"create,omitempty"`
	Read bool `json:"read,omitempty"`
}

func (e EdgeAnnotation) Name() string {
	return "rest"
}

type FieldAnnotation struct {
	Create bool `json:"create,omitempty"`
	Update bool `json:"update,omitempty"`
	Read bool `json:"read,omitempty"`
	QueryOps QueryOps `json:"queryOps,omitempty"`
}

type QueryOps interface {
	HasOp(name string) bool
}

func (f FieldAnnotation) Name() string {
	return "rest"
}

type BoolQueryOps struct {
	EQ bool `json:"EQ,omitempty"`
	NEQ bool `json:"NEQ,omitempty"`
}

func (o BoolQueryOps) HasOp(name string) bool {
	switch name {
	case "EQ":
		return o.EQ
	case "NEQ":
		return o.NEQ
	}
	return false
}

type EnumQueryOps struct {
	EQ bool `json:"EQ,omitempty"`
	NEQ bool `json:"NEQ,omitempty"`
	In bool `json:"In,omitempty"`
	NotIn bool `json:"NotIn,omitempty"`
}

func (o EnumQueryOps) HasOp(name string) bool {
	switch name {
	case "EQ":
		return o.EQ
	case "NEQ":
		return o.NEQ
	case "In":
		return o.In
	case "NotIn":
		return o.NotIn
	}
	return false
}

type NumericQueryOps struct {
	EQ bool `json:"EQ,omitempty"`
	NEQ bool `json:"NEQ,omitempty"`
	In bool `json:"In,omitempty"`
	NotIn bool `json:"NotIn,omitempty"`
	GT bool `json:"GT,omitempty"`
	GTE bool `json:"GTE,omitempty"`
	LT bool `json:"LT,omitempty"`
	LTE bool `json:"LTE,omitempty"`
}

func (o NumericQueryOps) HasOp(name string) bool {
	switch name {
	case "EQ":
		return o.EQ
	case "NEQ":
		return o.NEQ
	case "In":
		return o.In
	case "NotIn":
		return o.NotIn
	case "GT":
		return o.GT
	case "GTE":
		return o.GTE

	case "LT":
		return o.LT
	case "LTE":
		return o.LTE
	}
	return false
}

type StringQueryOps struct {
	EQ bool `json:"EQ,omitempty"`
	NEQ bool `json:"NEQ,omitempty"`
	In bool `json:"In,omitempty"`
	NotIn bool `json:"NotIn,omitempty"`
	GT bool `json:"GT,omitempty"`
	GTE bool `json:"GTE,omitempty"`
	LT bool `json:"LT,omitempty"`
	LTE bool `json:"LTE,omitempty"`
	Contains bool `json:"Contains,omitempty"`
	HasPrefix bool `json:"HasPrefix,omitempty"`
	HasSuffix bool `json:"HasSuffix,omitempty"`
}

func (o StringQueryOps) HasOp(name string) bool {
	switch name {
	case "EQ":
		return o.EQ
	case "NEQ":
		return o.NEQ
	case "In":
		return o.In
	case "NotIn":
		return o.NotIn
	case "GT":
		return o.GT
	case "GTE":
		return o.GTE
	case "LT":
		return o.LT
	case "LTE":
		return o.LTE
	case "Contains":
		return o.Contains
	case "HasPrefix":
		return o.HasPrefix
	case "HasSuffix":
		return o.HasSuffix
	}
	return false
}



