package vo

type OwnerMail string

func (v OwnerMail) Validate() bool {
	if v == "" {
		return false
	}
	// FIXME:
	return true
}

func (v OwnerMail) ToVal() string {
	return string(v)
}

func ParseOwnerMail(v string) OwnerMail {
	return OwnerMail(v)
}
