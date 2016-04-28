package controller

import (
	"github.com/myProject/model/dataModel"
	//"encoding/json"
	"fmt"
	"github.com/eaciit/dbox"
	_ "github.com/eaciit/dbox/dbc/json"
	"github.com/eaciit/knot/knot.v1"
	"github.com/eaciit/toolkit"
	"os"
	"path/filepath"
)

type UserController struct {
	App
}

type user struct {
	Nama string `json : "nama"`
	ID   string `json: "id"`
	Kota string `json: "kota"`
}

func CreateUserController(u *knot.Server) *UserController {
	var controller = new(UserController)
	controller.Server = u
	return controller
}

func connection() (dbox.IConnection, error) {
	//config := toolkit.M{"newfile": true}
	wd, _ := os.Getwd()
	ci := &dbox.ConnectionInfo{filepath.Join(wd, "model", "data", "coba.json"), "", "", "", nil}

	conn, err := dbox.NewConnection("json", ci)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	err = conn.Connect()
	return conn, nil

}

func (u *UserController) GetAll(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputJson
	conn, e := new(modelcore.User).Connection()
	if e != nil {
		return e
	}

	defer conn.Close()

	csr, e := conn.NewQuery().Cursor(nil)
	if e != nil {
		return e
	}

	defer csr.Close()

	result := []toolkit.M{}
	err := csr.Fetch(&result, 0, false)
	if err != nil {
		return nil
	}

	return result
}

func (u *UserController) GetSave(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputJson
	c, e := new(modelcore.User).Connection()
	if e != nil {
		fmt.Println(e.Error())
		return e
	}

	defer c.Close()
	query := c.NewQuery().Save()
	data := modelcore.User{}

	payload := toolkit.M{}
	err := r.GetPayload(&payload)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	//fmt.Println("-------- payload ", payload.data)
	//fmt.Println("-------- payload ", payload["data"])

	data.ID = payload["id"].(string)
	data.Nama = payload["nama"].(string)
	data.Kota = payload["kota"].(string)

	fmt.Println("----------- nama ", data.Nama)
	e = query.Exec(toolkit.M{
		"data": data,
	})
	var f = map[string]string{"data": "gagal"}

	if e != nil {
		fmt.Println("----------- error", e.Error())
		return f
	}

	//query.Close()

	a := map[string]string{"data": "berhasil"}

	return a
}

func (u *UserController) TestCoba(r *knot.WebContext) interface{} {
	data := new(modelcore.User).TestCoba()
	coba, _ := new(modelcore.User).Connection()
	fmt.Println("------------ koneksi", coba)
	return data
}
