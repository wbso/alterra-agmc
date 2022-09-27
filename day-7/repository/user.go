package repository

import (
	"alterraseven/entity"
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserDAO struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
}

func (u UserDAO) ToEntity() entity.User {
	return entity.User{
		ID:       int(u.ID),
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}

func usersDAOToEntities(users []UserDAO) (res []entity.User) {
	for _, user := range users {
		res = append(res, user.ToEntity())
	}
	return res
}

func userEntityToDAO(user entity.User) UserDAO {
	return UserDAO{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

type UserRepository interface {
	Index(ctx context.Context) ([]entity.User, error)
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	GetByID(ctx context.Context, id int) (entity.User, error)
	Update(ctx context.Context, user entity.User) (entity.User, error)
	Delete(ctx context.Context, user entity.User) error
	Create(ctx context.Context, user entity.User) (entity.User, error)
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u UserRepo) Index(ctx context.Context) ([]entity.User, error) {
	var users []UserDAO
	result := u.db.Find(&users)
	return usersDAOToEntities(users), result.Error
}

func (u UserRepo) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	var user UserDAO
	res := u.db.Where("email = ?", email).First(&user)
	if res.Error != nil {
		return user.ToEntity(), res.Error
	}

	return user.ToEntity(), nil
}

func (u UserRepo) GetByID(ctx context.Context, id int) (entity.User, error) {
	var user UserDAO
	res := u.db.First(&user, id)
	if res.Error != nil {
		return user.ToEntity(), res.Error
	}

	return user.ToEntity(), nil
}

func (u UserRepo) Update(ctx context.Context, data entity.User) (entity.User, error) {
	var user UserDAO
	u.db.Model(&user).
		Clauses(clause.Returning{}).
		Where("id = ?", data.ID).
		Update("name", data.Name)
	return user.ToEntity(), nil
}

func (u UserRepo) Delete(ctx context.Context, user entity.User) error {
	res := u.db.Delete(user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (u UserRepo) Create(ctx context.Context, user entity.User) (entity.User, error) {
	userDAO := userEntityToDAO(user)
	result := u.db.Create(&userDAO)
	return user, result.Error
}
