package usecase

type Guild interface {
	RequestCreateGuildByGuest()
}

func NewGuild() Guild {
	return &guild{}
}

type guild struct {
}

func (g *guild) RequestCreateGuildByGuest() {

}
