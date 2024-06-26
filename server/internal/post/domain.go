package post

import "time"

type Post struct {
	ID               int    `json:"id"`
	UUID             string `json:"uuid"`
	Content          string `json:"content"`
	NumLike          int64  `json:"num_like,omitempty"`
	UserId           int    `json:"user_id"`
	VisibilityTypeId int    `json:"visibility_type_id"`
	DeletedAt        string `json:"deleted_at"`
	UpdateAt         string `json:"update_at"`
}

type PostResponse struct {
	UUID             *string   `json:"uuid"`
	Content          *string   `json:"content"`
	NumLike          int64     `json:"num_like,omitempty"`
	Username         *string   `json:"username"`
	UserUUID         *string   `json:"user_uuid"`
	VisibilityTypeId int       `json:"visibility_type_id"`
	UpdateAt         time.Time `json:"update_at"`
}

type PostCreated struct {
	UUID             string `json:"uuid"`
	Content          string `json:"content"`
	NumLike          int64  `json:"num_like,omitempty"`
	UserUUID         string `json:"user_uuid"`
	VisibilityTypeId int    `json:"visibility_type_id"`
}

type VisibilityType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
