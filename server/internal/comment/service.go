package comment

type CommentService struct {
	commentRepo *CommentRepository
}

func NewcommentService(commRepo *CommentRepository) *CommentService {
	return &CommentService{commRepo}
}
