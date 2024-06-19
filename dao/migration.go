package dao

import "GuidingSystem/model/reply"

func (db *db) Migration() {
	if !db.DBEngine.HasTable(&reply.User{}) {
		db.DBEngine.CreateTable(&reply.User{})
	}
	if !db.DBEngine.HasTable(&reply.Cemetery{}) {
		db.DBEngine.CreateTable(&reply.Cemetery{})
	}
	if !db.DBEngine.HasTable(&reply.Attraction{}) {
		db.DBEngine.CreateTable(&reply.Attraction{})
	}
}
