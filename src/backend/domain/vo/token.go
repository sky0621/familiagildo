package vo

type Token string

func (v Token) Validate() error {
	// FIXME:
	return nil
}

func (v Token) FieldName() string {
	return "token"
}

func (v Token) ToVal() string {
	return string(v)
}

func ParseToken(v string) Token {
	return Token(v)
}
