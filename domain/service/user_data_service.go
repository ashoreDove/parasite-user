package service

import (
	"errors"
	"github.com/ashoreDove/parasite/user/domain/model"
	"github.com/ashoreDove/parasite/user/domain/repository"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//数据逻辑在这边做，业务逻辑在Handler
type IUserDataService interface {
	AddUser(*model.User) (string, error)
	UpdateUser(*model.User) error
	findUserByAccount(string) (*model.User, error)
	FindAllUser() ([]model.User, error)
	CheckPassword(string, string) (bool, error)
	generate(string) ([]byte, error)
	validate(string, string) (bool, error)
}

//创建
func NewUserDataService(db *gorm.DB) IUserDataService {
	return &UserDataService{repository.NewUserRepository(db)}
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}

func (u *UserDataService) CheckPassword(account string, pwd string) (bool, error) {
	user, err := u.findUserByAccount(account)
	if err != nil {
		return false, err
	}
	return u.validate(pwd, user.Password)
}

func (u *UserDataService) generate(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}

func (u *UserDataService) validate(pwd string, hashed string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pwd)); err != nil {
		return false, errors.New("密码错误")
	}
	return true, nil
}

func (u *UserDataService) findUserByAccount(account string) (*model.User, error) {
	return u.UserRepository.FindUserByAccount(account)
}

//插入
func (u *UserDataService) AddUser(user *model.User) (string, error) {
	pwdByte, err := u.generate(user.Password)
	if err != nil {
		return user.Account, err
	}
	user.Password = string(pwdByte)
	return u.UserRepository.CreateUser(user)
}

//更新
func (u *UserDataService) UpdateUser(user *model.User) error {
	pwdByte, err := u.generate(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(pwdByte)
	return u.UserRepository.UpdateUser(user)
}

//查找
func (u *UserDataService) FindAllUser() ([]model.User, error) {
	return u.UserRepository.FindAll()
}
