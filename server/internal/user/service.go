package user

type IUserRepository interface {
	GetUserByUUID(string) (User, error)
}

type UserService struct {
	userRepo IUserRepository
}

func NewUserService(userRepo IUserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (us *UserService) GetUserByUUID(s string) (User, error) {
	return us.userRepo.GetUserByUUID(s)
}
