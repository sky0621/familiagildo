// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Node interface {
	IsNode()
	GetID() string
}

type Task interface {
	IsTask()
}

// ギルドフィルター条件
type AdminGuildFilter struct {
	NameLike *string `json:"nameLike,omitempty"`
}

// ギルドインプット
type AdminGuildInput struct {
	// ギルドの名前
	Name string `json:"name"`
	// オーナーのメールアドレス
	OwnerMail string `json:"ownerMail"`
}

type CreateGuildByGuestInput struct {
	Token     string `json:"token"`
	OwnerName string `json:"ownerName"`
	LoginID   string `json:"loginID"`
	Password  string `json:"password"`
}

type CreateParticipantByGuestInput struct {
	LoginID  string `json:"loginID"`
	Password string `json:"password"`
}

// ゲストトークン
type GuestToken struct {
	// ID
	ID string `json:"id"`
	// オーナーメールアドレス
	Mail string `json:"mail"`
	// トークン
	Token string `json:"token"`
	// トークン有効期限
	ExpirationDate *time.Time `json:"expiration_date,omitempty"`
	// 受付番号
	AcceptedNumber string `json:"accepted_number"`
}

func (GuestToken) IsNode()            {}
func (this GuestToken) GetID() string { return this.ID }

type MutationResponse struct {
	ID *string `json:"id,omitempty"`
}

type NoopInput struct {
	ClientMutationID *string `json:"clientMutationId,omitempty"`
}

type NoopPayload struct {
	ClientMutationID *string `json:"clientMutationId,omitempty"`
}

// お知らせ
type Notice struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	SubTitle *string `json:"subTitle,omitempty"`
	Content  *string `json:"content,omitempty"`
	// 表示開始日時
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// 表示終了日時
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
}

type NoticeInput struct {
	Title    string  `json:"title"`
	SubTitle *string `json:"subTitle,omitempty"`
	Content  *string `json:"content,omitempty"`
	// 表示開始日時
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// 表示終了日時
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
}

type Owner struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Mail string `json:"mail"`
	// 所属ギルド
	Guild *Guild `json:"guild"`
	// 自分が登録したタスクリスト
	MyTasks []*OwnerTask `json:"myTasks"`
}

// オーナータスク
type OwnerTask struct {
	// ID
	ID string `json:"id"`
	// タスクの内容
	Content string `json:"content"`
	// タスクのステータス
	Status TaskStatus `json:"status"`
	// タスクの継続性
	Continuity TaskContinuity `json:"continuity"`
	// タスクの期日
	DueDateTime *time.Time `json:"dueDateTime,omitempty"`
}

func (OwnerTask) IsNode()            {}
func (this OwnerTask) GetID() string { return this.ID }

func (OwnerTask) IsTask() {}

// オーナータスクフィルター条件
type OwnerTaskFilter struct {
	ContentLike *string `json:"contentLike,omitempty"`
}

// オーナータスク
type OwnerTaskInput struct {
	// タスクの内容
	Content string `json:"content"`
	// タスクの期日
	DueDate *time.Time `json:"dueDate,omitempty"`
}

type Participant struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Mail *string `json:"mail,omitempty"`
	// 所属ギルド
	Guild *Guild `json:"guild"`
	// 自分が登録したタスクリスト
	MyTasks []*ParticipantTask `json:"myTasks"`
}

// 参加者タスク
type ParticipantTask struct {
	// ID
	ID string `json:"id"`
	// タスクの内容
	Content string `json:"content"`
	// タスクのステータス
	Status TaskStatus `json:"status"`
	// タスクの継続性
	Continuity TaskContinuity `json:"continuity"`
	// タスクの期日
	DueDateTime *time.Time `json:"dueDateTime,omitempty"`
}

func (ParticipantTask) IsNode()            {}
func (this ParticipantTask) GetID() string { return this.ID }

func (ParticipantTask) IsTask() {}

// 参加者タスクフィルター条件
type ParticipantTaskFilter struct {
	ContentLike *string `json:"contentLike,omitempty"`
}

// 参加者タスク
type ParticipantTaskInput struct {
	// タスクの内容
	Content string `json:"content"`
	// タスクの期日
	DueDate *time.Time `json:"dueDate,omitempty"`
}

type RequestCreateGuildInput struct {
	GuildName string `json:"guildName"`
	OwnerMail string `json:"ownerMail"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// タスクの継続性
type TaskContinuity string

const (
	// 単発
	TaskContinuityOneshot TaskContinuity = "ONESHOT"
	// 継続
	TaskContinuityContinuation TaskContinuity = "CONTINUATION"
)

var AllTaskContinuity = []TaskContinuity{
	TaskContinuityOneshot,
	TaskContinuityContinuation,
}

func (e TaskContinuity) IsValid() bool {
	switch e {
	case TaskContinuityOneshot, TaskContinuityContinuation:
		return true
	}
	return false
}

func (e TaskContinuity) String() string {
	return string(e)
}

func (e *TaskContinuity) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskContinuity(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskContinuity", str)
	}
	return nil
}

func (e TaskContinuity) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// タスクのステータス
type TaskStatus string

const (
	// 仮登録
	TaskStatusRegistering TaskStatus = "REGISTERING"
	// 登録済み
	TaskStatusRegistered TaskStatus = "REGISTERED"
	// 交渉中
	TaskStatusNegotiating TaskStatus = "NEGOTIATING"
	// 受諾済み
	TaskStatusAccepted TaskStatus = "ACCEPTED"
)

var AllTaskStatus = []TaskStatus{
	TaskStatusRegistering,
	TaskStatusRegistered,
	TaskStatusNegotiating,
	TaskStatusAccepted,
}

func (e TaskStatus) IsValid() bool {
	switch e {
	case TaskStatusRegistering, TaskStatusRegistered, TaskStatusNegotiating, TaskStatusAccepted:
		return true
	}
	return false
}

func (e TaskStatus) String() string {
	return string(e)
}

func (e *TaskStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskStatus", str)
	}
	return nil
}

func (e TaskStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}