package models

import (
	"database/sql"
	"github.com/lib/pq"
	"time"
	"html/template"
)

type Post struct {
	Id int `db:"id"`
	Board string `db:"board"`
	Picture template.URL `db:"picture"`
	Thread sql.NullInt32 `db:"thread"`
	ReplyTo sql.NullInt32 `db:"reply_to"`
	IpAddress sql.NullString `db:"ip_address"`
	Special sql.NullString `db:"special"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt pq.NullTime `db:"updated_at"`
}

type Thread struct {
	Id int `db:"id"`
	Board string `db:"board"`
	Picture template.URL `db:"picture"`
	Thread sql.NullInt32 `db:"thread"`
	ReplyTo sql.NullInt32 `db:"reply_to"`
	IpAddress sql.NullString `db:"ip_address"`
	Special sql.NullString `db:"special"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt pq.NullTime `db:"updated_at"`
	Replies int `db:"replies"`
}
