package model

import (
	"fmt"
	"time"
)

type User struct {
	Id       int
	Name     string `sql:"not null"`
	Password string `sql:"not null"`
	Ctime    time.Time
}

func (u *User) String() string {
	return fmt.Sprintf("%#v", *u)
}

func (u *User) Get() {
	// Db.QueryRow(`SELECT name,password,ctime
	// FROM "user"
	// WHERE id=$1`, u.Id).Scan(&u.Name, &u.Password, &u.Ctime)
	Db.First(u, u.Id)
}

func (u *User) GetByName() {
	// row := Db.QueryRow(`SELECT id,password,ctime
	// FROM "user"
	// WHERE name=$1`, u.Name)
	// fmt.Println(*row)
	// row.Scan(&u.Id, &u.Password, &u.Ctime)
	Db.Where("name=?", u.Name).First(u)
}

func (u *User) CheckPassword(pwd string) bool {
	return u.Password == pwd
}

func CheckLogin(uname string, pwd string) (u User, ok bool) {
	u = User{Name: uname}
	u.GetByName()
	// fmt.Println(u.String())
	if u.Id == 0 {
		ok = false
		return
	}
	ok = u.CheckPassword(pwd)
	return
}
