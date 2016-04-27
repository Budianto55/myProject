package controller

import (
	"fmt"
	"github.com/eaciit/knot/knot.v1"
	"os"
)

type App struct {
	Server *knot.Server
}

var (
	LayoutFile   string   = "views/layout.html"
	IncludeFiles []string = []string{"views/_head.html"}
	AppBasepath  string   = func(dir string, err error) string { return dir }(os.Getwd())
)

func init() {
	fmt.Println("Base Path ====> ", AppBasepath)
}
