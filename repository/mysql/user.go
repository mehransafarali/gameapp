package mysql

import (
	"GameApp/entity"
	"GameApp/service/userservice"
	"database/sql"
	"fmt"
)

func (d MySQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	user := entity.User{}
	var createdAt []uint8
	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, fmt.Errorf("mysql query error: %w", err)
	}

	return false, nil
}

func (d MySQLDB) Register(u userservice.RegisterRequest) (entity.User, error) {
	res, err := d.db.Exec(`insert into users(name, phone_number, password) values(?, ?, ?)`, u.Name, u.PhoneNumber, u.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("can not insert user: %w", err)
	}

	id, _ := res.LastInsertId()
	var RegisteredUser entity.User
	RegisteredUser.ID = uint(id)
	return RegisteredUser, nil
}

func (d MySQLDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error) {
	user := entity.User{}
	var createdAt []uint8
	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, false, nil
		}
		return entity.User{}, false, fmt.Errorf("mysql query error: %w", err)
	}

	return user, true, nil
}

func (d MySQLDB) Login(phone_number, password string) (bool, error) {
	row := d.db.QueryRow(`select phone_number, password from users where phone_number =?  and password= ?`, phone_number, password)
	var resPhoneNumber, resPassword string
	err := row.Scan(&resPhoneNumber, &resPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("mysql query error: %w", err)
	}

	return true, nil
}

func (d MySQLDB) GetUserByID(id uint) (entity.User, error) {
	user := entity.User{}
	var createdAt []uint8
	row := d.db.QueryRow(`select * from users where id = ?`, id)
	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, fmt.Errorf("user not found")
		}
		return entity.User{}, fmt.Errorf("mysql query error: %w", err)
	}

	return user, nil
}
