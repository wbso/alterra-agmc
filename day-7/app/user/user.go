package user

import (
	"alterraseven/dto"
	"alterraseven/entity"
	"alterraseven/repository"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Index(context.Context) ([]dto.UserResponse, error)
	Get(context.Context, int) (dto.UserResponse, error)
	Create(context.Context, dto.UserCreateRequest) (dto.UserResponse, error)
	Update(context.Context, int, dto.UserCreateRequest) (dto.UserResponse, error)
	Delete(context.Context, int) error
}

type UserService struct {
	repo repository.UserRepository
}

func New(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (u UserService) Index(ctx context.Context) (res []dto.UserResponse, err error) {
	index, err := u.repo.Index(ctx)
	if err != nil {
		return res, err
	}
	return userEntitiesToDTO(index), nil
}

func (u UserService) Get(ctx context.Context, i int) (res dto.UserResponse, err error) {
	user, err := u.repo.GetByID(ctx, i)
	if err != nil {
		return res, err
	}
	return userEntityToDTO(user), nil
}

func (u UserService) Create(ctx context.Context, input dto.UserCreateRequest) (res dto.UserResponse, err error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return res, err
	}
	createdUser, err := u.repo.Create(ctx, entity.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hash,
	})
	if err != nil {
		return res, err
	}

	return userEntityToDTO(createdUser), nil
}

func (u UserService) Update(ctx context.Context, i int, input dto.UserCreateRequest) (res dto.UserResponse, err error) {
	//
	//	// validate authorization
	//	if id != authId {
	//		return c.String(http.StatusForbidden, "you dont have permission to access this resource")
	//	}
	//
	updatedUser, err := u.repo.Update(ctx, entity.User{
		ID:    i,
		Name:  input.Name,
		Email: input.Email,
	})
	if err != nil {
		return dto.UserResponse{}, err
	}
	return userEntityToDTO(updatedUser), err

}

func (u UserService) Delete(ctx context.Context, i int) error {
	err := u.repo.Delete(ctx, entity.User{
		ID: i,
	})
	if err != nil {
		return err
	}
	return nil
}

func userEntityToDTO(user entity.User) dto.UserResponse {
	return dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func userEntitiesToDTO(users []entity.User) (res []dto.UserResponse) {
	for _, user := range users {
		res = append(res, userEntityToDTO(user))
	}
	return res
}
