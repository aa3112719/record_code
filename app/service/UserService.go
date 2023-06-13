package service

import (
	"record_code/app/model"
	"record_code/utils/psql"
)

func GetDetailById(id string) model.User {
	var users = model.User{}
	psql.DB.Table("user").Where("id = ?", id).First(&users)

	return users
}
