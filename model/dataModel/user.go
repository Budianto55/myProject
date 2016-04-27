package modelcore

import (
	"fmt"
	// "github.com/eaciit/orm/v1"
	// "github.com/eaciit/toolkit"
	"github.com/eaciit/dbox"
	_ "github.com/eaciit/dbox/dbc/json"
	"os"
	"path/filepath"
)

type User struct {
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	Kota   string `json:"kota"`
}

// func (a *User) TableName() string {
// 	return user

// }

// func (a *User) RecordID() interface{} {
// 	return a.Alamat
// }

func (u *User) TestCoba() interface{} {
	var a = "nanananaanana"
	return a
}

func (u *User) Connection() (dbox.IConnection, error) {
	wd, _ := os.Getwd()
	ci := &dbox.ConnectionInfo{filepath.Join(wd, "model", "data", "coba.json"), "", "", "", nil}

	conn, err := dbox.NewConnection("json", ci)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = conn.Connect()
	return conn, nil
}
