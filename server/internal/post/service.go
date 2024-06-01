package post

type PostService struct {
	postRepo *PostRepository
}

func NewPostService(postRepo *PostRepository) *PostService {
	return &PostService{postRepo}
}
