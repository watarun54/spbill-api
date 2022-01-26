package usecase

import "github.com/watarun54/serverless-skill-manager/server/domain"

type (
	IUserRepository interface {
		FindById(id int) (domain.User, error)
		FindByIds(ids []int) (domain.Users, error)
		FindByEmail(email string) (domain.User, error)
		FindByLineID(lineID string) (domain.User, error)
		FindAll() (domain.Users, error)
		Store(domain.User) (domain.User, error)
		Update(domain.User) (domain.User, error)
		DeleteById(domain.User) error
	}

	UserInteractor struct {
		UserRepository IUserRepository
	}
)

func (interactor *UserInteractor) ConvertUserFormToUser(userForm domain.UserForm) (user domain.User) {
	user.ID = userForm.ID
	user.Name = userForm.Name
	user.Email = userForm.Email
	user.HashedPassword = userForm.HashedPassword
	user.LineID = userForm.LineID
	return
}

func (interactor *UserInteractor) UserById(id int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindById(id)
	return
}

func (interactor *UserInteractor) UserByEmail(email string) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindByEmail(email)
	return
}

func (interactor *UserInteractor) UserByLineID(lineID string) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindByLineID(lineID)
	return
}

func (interactor *UserInteractor) Users() (users domain.Users, err error) {
	users, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) Add(u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.Store(u)
	return
}

func (interactor *UserInteractor) Update(u domain.User) (user domain.User, err error) {
	user, err = interactor.UserRepository.Update(u)
	return
}

func (interactor *UserInteractor) DeleteById(u domain.User) (err error) {
	err = interactor.UserRepository.DeleteById(u)
	return
}
