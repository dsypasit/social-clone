package comment

type comment struct {
	ID        int    `json:"id"`
	UUID      string `json:"uuid"`
	Content   string `json:"content"`
	UserId    int    `json:"user_id"`
	PostId    int    `json:"post_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
