package repository

import (
	log "github.com/sirupsen/logrus"
	"go-gin-gorm-wire/app/domain/dao"
	"gorm.io/gorm"
)

type RoleRepository interface {
	FindAllRole() ([]dao.Role, error)
	FindRoleById(id int) (dao.Role, error)
	Save(Role *dao.Role) (dao.Role, error)
	DeleteRoleById(id int) error
}

type RoleRepositoryImpl struct {
	db *gorm.DB
}

func (u RoleRepositoryImpl) FindAllRole() ([]dao.Role, error) {
	var Roles []dao.Role

	var err = u.db.Find(&Roles).Error
	if err != nil {
		log.Error("Got an error finding all couples. Error: ", err)
		return nil, err
	}

	return Roles, nil
}

func (u RoleRepositoryImpl) FindRoleById(id int) (dao.Role, error) {
	Role := dao.Role{
		ID: id,
	}
	err := u.db.First(&Role).Error
	if err != nil {
		log.Error("Got and error when find Role by id. Error: ", err)
		return dao.Role{}, err
	}
	return Role, nil
}

func (u RoleRepositoryImpl) Save(Role *dao.Role) (dao.Role, error) {
	var err = u.db.Save(Role).Error
	if err != nil {
		log.Error("Got an error when save Role. Error: ", err)
		return dao.Role{}, err
	}
	return *Role, nil
}

func (u RoleRepositoryImpl) DeleteRoleById(id int) error {
	err := u.db.Delete(&dao.Role{}, id).Error
	if err != nil {
		log.Error("Got an error when delete Role. Error: ", err)
		return err
	}
	return nil
}

func RoleRepositoryInit(db *gorm.DB) *RoleRepositoryImpl {
	db.AutoMigrate(&dao.Role{})
	return &RoleRepositoryImpl{
		db: db,
	}
}
