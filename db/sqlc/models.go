// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"database/sql"
)

type Event struct {
	ID               int64          `json:"id"`
	UserID           int64          `json:"user_id"`
	Name             string         `json:"name"`
	IpAddress        sql.NullString `json:"ip_address"`
	UserAgent        sql.NullString `json:"user_agent"`
	CreatedBy        sql.NullInt64  `json:"created_by"`
	UpdatedBy        sql.NullInt64  `json:"updated_by"`
	CreatedTimestamp sql.NullTime   `json:"created_timestamp"`
	UpdatedTimestamp sql.NullTime   `json:"updated_timestamp"`
}

type LoginDetail struct {
	ID               int64         `json:"id"`
	UserID           int64         `json:"user_id"`
	PasswordHash     string        `json:"password_hash"`
	IsLockedOut      sql.NullBool  `json:"is_locked_out"`
	FailedLoginCount sql.NullInt32 `json:"failed_login_count"`
	CreatedBy        sql.NullInt64 `json:"created_by"`
	UpdatedBy        sql.NullInt64 `json:"updated_by"`
	CreatedTimestamp sql.NullTime  `json:"created_timestamp"`
	UpdatedTimestamp sql.NullTime  `json:"updated_timestamp"`
}

type User struct {
	ID               int64          `json:"id"`
	FirstName        sql.NullString `json:"first_name"`
	LastName         sql.NullString `json:"last_name"`
	FullName         sql.NullString `json:"full_name"`
	IsEmailConfirmed sql.NullBool   `json:"is_email_confirmed"`
	Email            string         `json:"email"`
	UserName         string         `json:"user_name"`
	RoleID           sql.NullInt32  `json:"role_id"`
	CreatedBy        sql.NullInt64  `json:"created_by"`
	UpdatedBy        sql.NullInt64  `json:"updated_by"`
	CreatedTimestamp sql.NullTime   `json:"created_timestamp"`
	UpdatedTimestamp sql.NullTime   `json:"updated_timestamp"`
	IsArchived       sql.NullBool   `json:"is_archived"`
}
