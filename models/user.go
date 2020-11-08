package models

import (
	"BeegoDemo/db_mysql"
	"BeegoDemo/util"
	"fmt"
)

type User struct {
	Id			int			`form:"id"`
	Phone 		string		`form:"phone"`
	Password	string		`form:"password"`
	Name 		string		`form:"name"`
	Card 		string		`form:"card"`
	Sex 		string		`form:"sex"`
}

func (u User) Update() (int64, error) {
	rs, err := db_mysql.Db.Exec("update user set name=?, card=?, sex=? where phone=?", u.Name, u.Card, u.Sex, &u.Phone)
	if err != nil {
		return -1, err
	}
	return rs.RowsAffected()
}


/**
 * 保存用户信息的方法：保存用户信息到数据中
 */
func (u User) SaveUser()(int64,error){
	//1、密码脱敏处理（hash加密）
	u.Password = util.MD5HashString(u.Password)
	//执行数据库操作
	fmt.Println("将要保存的手机号码：", u.Phone, "密码：", u.Password)
	result, err := db_mysql.Db.Exec("insert into user(phone, password) " +
		"values(?,?)", u.Phone, u.Password)
	if err != nil { //保存数据时遇到错误
		return -1, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, nil
}

/**
 * 查询用户信息
 */
func (u User) QueryUser() (*User, error) {
	//1、密码脱敏处理
	u.Password = util.MD5HashString(u.Password)

	row := db_mysql.Db.QueryRow("select phone, name, card, sex from user where  phone = ? and password = ?",
		u.Phone, u.Password)
	err := row.Scan(&u.Phone, &u.Name, &u.Card, &u.Sex)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
/**
 * 根据用户的phone信息查询对应的用户的实名信息
 */
func QueryUserByPhone(phone string) (*User, error) {
	row := db_mysql.Db.QueryRow("select phone, name, card, sex from user where phone = ?", phone)
	var user User
	err := row.Scan(&user.Phone, &user.Name, &user.Card, &user.Sex)
	if err != nil {
		return nil, err
	}
	return &user, nil
}