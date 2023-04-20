package vo

type Token string

func (v Token) Validate() bool {
	if v == "" {
		return false
	}
	// FIXME:
	return true
}

func (v Token) ToVal() string {
	return string(v)
}

func ParseToken(v string) Token {
	return Token(v)
}
