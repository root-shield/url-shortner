package shorturl

import (
	"database/sql"
)

type ShortUrl struct {
	Id                 int64          `json:"id"`
	Url                string         `json:"url"`
	UrlHash            string         `json:"hash"`
	HashInBase32       string         `json:"hash_in_base32"`
	ShortBase32        string         `json:"short_base32"`
	ShortBase32Shuffle bool           `json:"short_base32_shuffle"`
	UrlStatus          sql.NullInt64  `json:"url_status"`
	LastCheckTime      sql.NullString `json:"last_check_time"`
	RedirectCount      int64          `json:"redirect_count"`
}

type ShortUrlRequest struct {
	Url         string `json:"url"`
	UrlHash     string `json:"hash"`
	ShortBase32 string `json:"short_base32"`
}

func (s *ShortUrlRequest) ValidateUrl() {
	//TODO валидация ссылки?
}
