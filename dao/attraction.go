package dao

import (
	"GuidingSystem/model/reply"
)

func (db *db) GetAttractionsByCemeteryID(cemeteryID int64) ([]*reply.Attraction, error) {
	res := make([]*reply.Attraction, 0)
	err := db.DBEngine.Where("cemetery_id = ?", cemeteryID).Find(&res).Error
	return res, err
}
