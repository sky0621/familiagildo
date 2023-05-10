package model

// Guild is ギルド
type Guild struct {
	// ID
	ID GuildID `json:"id"`
	// ギルドの名称
	Name string `json:"name"`
	// オーナー
	Owner *Owner `json:"owner"`
	// 参加者リスト
	Participants []*Participant `json:"participants"`
	// タスクリスト
	Tasks []Task `json:"tasks"`
}

func (Guild) IsNode()         {}
func (g Guild) GetID() string { return g.ID.NodeID() }
