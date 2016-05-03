package modelcore

import (
	"fmt"
	"github.com/eaciit/dbox"
	_ "github.com/eaciit/dbox/dbc/mongo"
	"github.com/eaciit/toolkit"
	"os"
	"path/filepath"
)

type User struct {
	Nama string `json:"nama", bson:"nama"`
	ID   string `json:"id", bson: "id"`
	Kota string `json:"kota", bson:"kota"`
}

type Member struct {
	ID     string `json:"_id", bson: "_id"`
	Nama   string `json:"nama", bson:"nama"`
	Alamat string `json:"alamat", bson:"alamat"`
	Kota   string `json:"kota", bson:"kota"`
}

func (a *User) TestCoba() interface{} {
	var m = "nanananaanana"
	return m
}

func (a *User) Connection() (dbox.IConnection, error) {
	wd, _ := os.Getwd()
	ci := &dbox.ConnectionInfo{filepath.Join(wd, "model", "data", "coba.json"), "", "", "", nil}

	conn, err := dbox.NewConnection("json", ci)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = conn.Connect()
	return conn, nil
}

func (m *Member) MemberConnect() (dbox.IConnection, error) {
	fmt.Println("----------------- masuk ")
	var conf = toolkit.M{}.Set("timeout", 3)
	ci := &dbox.ConnectionInfo{"localhost:27017", "UserDB", "", "", conf}
	conn, e := dbox.NewConnection("mongo", ci)
	if e != nil {
		fmt.Println("error -------", e.Error())
		return nil, e
	}

	e = conn.Connect()
	if e != nil {
		return nil, e
	}
	return conn, nil
}
