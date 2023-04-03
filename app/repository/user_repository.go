package repository

import (
	log "github.com/sirupsen/logrus"
	"go-gin-gorm-wire/app/domain/dao"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAllUser() ([]dao.User, error)
	FindUserById(id int) (dao.User, error)
	Save(User *dao.User) (dao.User, error)
	DeleteUserById(id int) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (u UserRepositoryImpl) FindAllUser() ([]dao.User, error) {
	var Users []dao.User

	var err = u.db.Preload("Role").Find(&Users).Error
	if err != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	return Users, nil
}

func (u UserRepositoryImpl) FindUserById(id int) (dao.User, error) {
	User := dao.User{
		ID: id,
	}
	err := u.db.Preload("Role").First(&User).Error
	if err != nil {
		log.Error("Got and error when find User by id. Error: ", err)
		return dao.User{}, err
	}
	return User, nil
}

func (u UserRepositoryImpl) Save(User *dao.User) (dao.User, error) {
	var err = u.db.Save(User).Error
	if err != nil {
		log.Error("Got an error when save User. Error: ", err)
		return dao.User{}, err
	}
	return *User, nil
}

func (u UserRepositoryImpl) DeleteUserById(id int) error {
	err := u.db.Delete(&dao.User{}, id).Error
	if err != nil {
		log.Error("Got an error when delete User. Error: ", err)
		return err
	}
	return nil
}

func UserRepositoryInit(db *gorm.DB) *UserRepositoryImpl {
	db.AutoMigrate(&dao.User{})
	return &UserRepositoryImpl{
		db: db,
	}
}
