package routes

import (
	//5.3
	"go-workshop/controllers"
	c "go-workshop/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func TestRoutes(app *fiber.App) {

	//5.0
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"gofiber": "21022566",
		},
	}))

	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/", c.HelloPokemon)
	//5.1
	v1.Get("/fact/:num", c.FactorialNumber)

	//6
	v1.Post("/register", c.TestUser)

	//CRUD dogs
	dog := v1.Group("/dog")
	dog.Get("", controllers.GetDogs)
	dog.Get("/", controllers.GetDogs)
	dog.Get("/deleted", controllers.GetDeletedDogs)
	dog.Get("/filter", controllers.GetDog)
	dog.Get("/json", controllers.GetDogsJson)
	dog.Post("/", controllers.AddDog) // create
	dog.Put("/:id", controllers.UpdateDog)
	dog.Delete("/:id", controllers.RemoveDog)

	dog.Get("/len", controllers.GetDogsLen)

	dog.Get("/between", controllers.GetBetween)

	dog.Get("/jsonv2", controllers.GetDogsJsonV2)

	//CRUD company
	cpn := v1.Group("/company")
	cpn.Get("", controllers.GetCompanies)
	cpn.Get("/filter", controllers.GetCompany)
	cpn.Post("/", controllers.AddCompany)
	cpn.Put("/:id", controllers.UpdateCompany)
	cpn.Delete("/:id", controllers.RemoveCompany)

	//5.2
	v3 := api.Group("/v3")
	v3.Post("/earn", c.QueryTest)

}
