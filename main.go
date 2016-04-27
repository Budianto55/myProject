package main

import (
	"github.com/eaciit/colony-core/v0"
	"github.com/eaciit/knot/knot.v1"
	"github.com/myProject/controller"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

var (
	server *knot.Server
)

func main() {
	runtime.GOMAXPROCS(4)

	wd, _ := os.Getwd()
	colonycore.ConfigPath = filepath.Join(wd, "config")

	knot.SharedObject().Set("FilePath", path.Join(controller.AppBasepath, "config", "files"))

	server = new(knot.Server)
	server.Address = "localhost:4444"
	server.RouteStatic("res", path.Join(controller.AppBasepath, "assets"))
	server.Register(controller.CreateWebController(server), "")
	server.Register(controller.CreateUserController(server), "")

	server.Route("/", func(r *knot.WebContext) interface{} {
		http.Redirect(r.Writer, r.Request, "/web/index", 301)
		return true
	})
	server.Listen()

}
