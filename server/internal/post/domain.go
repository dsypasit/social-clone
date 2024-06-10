package post

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
	ID               int    `json:"id"`
	UUID             string `json:"uuid"`
	Content          string `json:"content"`
	NumLike          int64  `json:"num_like,omitempty"`
	UserUUID         string `json:"user_uuid"`
	VisibilityTypeId int    `json:"visibility_type_id"`
	DeletedAt        string `json:"deleted_at"`
	UpdateAt         string `json:"update_at"`
}

type VisibilityType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
