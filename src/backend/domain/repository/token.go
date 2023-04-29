package repository

type GuestTokenRepository interface {
	//
	GetByMailWithinValidPeriod()
}
