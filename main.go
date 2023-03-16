package main

import (
	"github.com/TugasEksplorasi/controllers"
	"github.com/go-martini/martini"
)

func main() {
	// Membuat instance Martini
	m := martini.Classic()

	// Routes
	r := martini.NewRouter()
	r.Get("/users/:id", controllers.GetUser)
	r.Post("/users", controllers.CreateUser)
	r.Put("/users/:id", controllers.UpdateUser)
	r.Delete("/users/:id", controllers.DeleteUser)

	// Hubungkan router ke aplikasi Martini
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)

	// Menjalankan aplikasi pada port 8000
	m.RunOnAddr(":8080")
}
