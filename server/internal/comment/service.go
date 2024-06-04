package comment

type commentService struct {
	commentRepo *commentRepository
}

func NewcommentService(commRepo *commentRepository) *commentService {
	return &commentService{commRepo}
}
