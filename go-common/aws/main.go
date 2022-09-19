package aws

type AWS struct {
	env string
}

func NewAWS(e string) *AWS {
	return &AWS{
		env: e,
	}
}
