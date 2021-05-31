package domain

type Params struct {
	Name   string
	Region string
}

type Flags struct {
	Name   string
	Region string
	Skip   int
	Limit  int
}
