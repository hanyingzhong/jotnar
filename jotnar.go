package jotnar

type Jotnar struct{}

func New() *Jotnar {
	return new(Jotnar)
}

// customer common initialization
func (*Jotnar) Init(functions ...func()) {
	for _, f := range functions {
		if f != nil {
			f()
		}
	}
}

func emptyStr(s string) bool {
	return s == ""
}
