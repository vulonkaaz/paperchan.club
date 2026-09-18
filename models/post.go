package models

import (
	"database/sql"
	"html/template"
	"time"
)

type Post struct {
	Id int `db:"id"`
	Board string `db:"board"`
	Picture template.URL `db:"picture"`
	Thread sql.NullInt32 `db:"thread"`
	ReplyTo sql.NullInt32 `db:"reply_to"`
	IpAddress sql.NullString `db:"ip_address"`
	Special sql.NullString `db:"special"`
	CreatedAt int64 `db:"created_at"`
	CreatedAtFormatted time.Time  // this isn't stored in the database but generated from CreatedAt, it mostly exist for display
}

type Thread struct {
	Id int `db:"id"`
	Board string `db:"board"`
	Picture template.URL `db:"picture"`
	Thread sql.NullInt32 `db:"thread"`
	ReplyTo sql.NullInt32 `db:"reply_to"`
	IpAddress sql.NullString `db:"ip_address"`
	Special sql.NullString `db:"special"`
	CreatedAt int64 `db:"created_at"`
	CreatedAtFormatted time.Time  // this isn't stored in the database but generated from CreatedAt, it mostly exist for display
	Replies int `db:"replies"`
}
