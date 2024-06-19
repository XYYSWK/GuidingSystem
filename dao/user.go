package dao

import (
	"GuidingSystem/model/reply"
)

func (db *db) CreateUser(user *reply.User) (reply.User, error) {
	res := reply.User{}
	result := db.DBEngine.Create(user)
	if result.Error != nil {
		return res, result.Error
	}
	result.Find(&res)
	return res, nil
}

func (db *db) GetUserByName(name string) (reply.User, error) {
	user := reply.User{}
	err := db.DBEngine.Where("name = ?", name).First(&user).Error
	return user, err
}

func (db *db) UpdateUser(user *reply.User) error {
	err := db.DBEngine.Save(user).Error
	return err
}
