package crypto

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/mutecomm/go-sqlcipher/v4"
)

func Opendb(path, password string) (*sql.DB, error) {
	escapedPassword := url.QueryEscape(password)
	dsn := fmt.Sprintf("%s?_pragma_key=%s&_pragma_cipher_page_size=4096", path, escapedPassword)

	return sql.Open("sqlite3", dsn)
}
