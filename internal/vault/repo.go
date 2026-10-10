package vault

import (
	"database/sql"
)

type PasswordRepo struct {
	db *sql.DB
}

func NewPasswordRepo(db *sql.DB) (*PasswordRepo, error) {
	r := PasswordRepo{db: db}
	err := r.Init()
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (r *PasswordRepo) Init() error {
	_, err := r.db.Exec(`CREATE TABLE IF NOT EXISTS Passwords (
		id          INTEGER PRIMARY KEY AUTOINCREMENT,
		service_name TEXT,
		url         TEXT,
		login       TEXT,
		password    TEXT
	)`)

	return err
}

func (r *PasswordRepo) DeleteEntry(id int) error {
	_, err := r.db.Exec("DELETE FROM Passwords WHERE id = ?", id)
	return err
}

func (r *PasswordRepo) WriteEntry(e Entry) error {
	_, err := r.db.Exec(`INSERT INTO Passwords (
			service_name, 
			url, 
			login, 
			password
		) VALUES (?, ?, ?, ?)`,
		e.ServiceName,
		e.URL,
		e.Login,
		e.Password,
	)

	return err
}

func (r *PasswordRepo) ChangeEntry(e Entry) error {
	_, err := r.db.Exec(`UPDATE Passwords SET
		service_name = ?,
		url = ?,
		login = ?,
		password = ?,
		WHERE id = ?`,
		e.ServiceName,
		e.URL,
		e.Login,
		e.Password,
		e.ID,
	)

	return err
}

func (r *PasswordRepo) GetEntries() ([]*Entry, error) {
	rows, err := r.db.Query("SELECT * FROM Passwords")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	entries := []*Entry{}
	for rows.Next() {
		e := &Entry{}
		err = rows.Scan(&e.ID, &e.ServiceName, &e.URL, &e.Login, &e.Password)
		if err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}
	return entries, nil
}
