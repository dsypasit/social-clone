package commend

type CommendService struct {
	commendRepo *CommendRepository
}

func NewCommendService(commRepo *CommendRepository) *CommendService {
	return &CommendService{commRepo}
}
