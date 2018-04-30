package main

import (
	"github.com/GinEduTren/controller"
	_ "github.com/lib/pq"
)

// Pull Main Controller
func main() {
	AllController()
}

// Sub Controller untuk membagi jenis-jenis controllernya
func AllController() {
	controller.MainController()
}
