package dao

import (
	"GuidingSystem/model/reply"
	"GuidingSystem/model/request"
)

func (db *db) CreateCemetery(cemetery *request.Cemetery) (reply.Cemetery, error) {
	res := reply.Cemetery{}
	result := db.DBEngine.Create(cemetery)
	if result.Error != nil {
		return res, result.Error
	}
	db.DBEngine.Find(&res)
	return res, nil
}

func (db *db) GetCemeteryByID(id int64) (reply.Cemetery, error) {
	user := reply.Cemetery{}
	err := db.DBEngine.Where("id = ?", id).FirstOrInit(&user).Error
	return user, err
}

func (db *db) GetCemeteries() ([]*reply.Cemetery, error) {
	res := make([]*reply.Cemetery, 0)
	err := db.DBEngine.Find(&res).Error
	return res, err
}

func (db *db) GetCemeteriesByName(name string) ([]*reply.Cemetery, error) {
	res := make([]*reply.Cemetery, 0)
	err := db.DBEngine.Where("name LIKE ?", "%"+name+"%").Find(&res).Error
	return res, err
}
