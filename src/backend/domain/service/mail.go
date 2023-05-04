package service

import "time"

const MailFormat = "1月2日 PM3時04分"

func ToMailFormattedDatetime(dt time.Time) string {
	return dt.Format(MailFormat)
}
