package vo

type TaskID string

func (v TaskID) Validate() bool {
	if v == "" {
		return false
	}
	return true
}
