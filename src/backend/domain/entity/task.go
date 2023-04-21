package entity

// Task ... 保護者が子どもにやってもらいたいこと、ないし、子どもが（対価を望む）こと
type Task struct {
	// ID ... タスクをユニークに特定するID
	//	ID vo.TaskID

	// Content ... タスクの内容
	Content string

	// CreatedBy ... タスクの作成者
	CreatedBy string
}
