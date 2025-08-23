package mysql

import "GameApp/entity"

//func (d DB)  {}

func (d DB) isPhoneNumberUnique(phoneNumber string) (bool, error) { return false, nil }

func (d DB) Register(u entity.User) (entity.User, error) { return entity.User{}, nil }
