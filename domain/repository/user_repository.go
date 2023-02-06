package repository

import (
	"github.com/ashoreDove/parasite-user/domain/model"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	InitTable() error
	FindUserByAccount(string) (*model.User, error)
	CreateUser(*model.User) (string, error)
	DeleteUserByAccount(string) error
	UpdateUser(*model.User) error
	FindAll() ([]model.User, error)
}

//创建userRepository
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDb: db}
}

type UserRepository struct {
	mysqlDb *gorm.DB
}

//初始化表
func (u *UserRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.User{}).Error
}

//根据ID查找User信息
func (u *UserRepository) FindUserByAccount(account string) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.First(user, account).Error
}

//创建User信息
func (u *UserRepository) CreateUser(user *model.User) (string, error) {
	return user.Account, u.mysqlDb.Create(user).Error
}

//根据ID删除User信息
func (u *UserRepository) DeleteUserByAccount(account string) error {
	return u.mysqlDb.Where("account = ?", account).Delete(&model.User{}).Error
}

//更新User信息
func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDb.Model(user).Update(user).Error
}

//获取结果集
func (u *UserRepository) FindAll() (userAll []model.User, err error) {
	return userAll, u.mysqlDb.Find(&userAll).Error
}
