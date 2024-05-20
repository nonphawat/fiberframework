package controllers

import (
	"strconv"
	"strings"

	"go-workshop/database"
	m "go-workshop/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HelloPokemon(c *fiber.Ctx) error {
	return c.SendString("Hello Pokemon and Pikachu")
}

func FactorialNumber(c *fiber.Ctx) error {
	getNum := c.Params("num")
	marks, err := strconv.Atoi(getNum)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	ttl := 1
	for i := 1; i <= marks; i++ {
		ttl *= i
	}

	return c.JSON(fiber.Map{
		"factorialNum": ttl,
	})

}

func QueryTest(c *fiber.Ctx) error {
	a := c.Query("tax_id")
	asciiValues := make([]int, len(a))
	for i, char := range a {
		asciiValues[i] = int(char)
	}
	return c.JSON(asciiValues)
}

func TestUser(c *fiber.Ctx) error {

	user := new(m.User)

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	validate := validator.New()
	errors := validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}

	return c.JSON(fiber.Map{
		"alldetailsUser": user,
	})
}

func DogIDGreaterThan100(db *gorm.DB) *gorm.DB {
	return db.Where("dog_id > ?", 100)
}

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Scopes(DogIDGreaterThan100).Find(&dogs)
	//db.Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

func GetDogsLen(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //delelete = null
	//db.Find(&dogs) //delelete = null
	return c.Status(200).JSON(len(dogs))
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func AddDog(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(fiber.Map{
		"dogs": dog,
	})

}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //10ตัว

	var dataResults []m.DogsRes
	for _, v := range dogs { //1 inet 112 //2 inet1 113
		typeStr := ""
		if v.DogID == 111 {
			typeStr = "red"
		} else if v.DogID == 113 {
			typeStr = "green"
		} else if v.DogID == 999 {
			typeStr = "pink"
		} else {
			typeStr = "no color"
		}

		d := m.DogsRes{
			Name:  v.Name,  //inet1
			DogID: v.DogID, //113
			Type:  typeStr, //green
		}
		dataResults = append(dataResults, d)
		// sumAmount += v.Amount
	}

	r := m.ResultData{
		Data:  dataResults,
		Name:  "golang-test",
		Count: len(dogs), //หาผลรวม,
	}
	return c.Status(200).JSON(r)
}

func GetDeletedDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dog []m.Dogs

	// will retrieve only the soft deleted records
	db.Unscoped().Where("deleted_at IS NOT NULL").Find(&dog)

	return c.Status(200).JSON(dog)
}

func DogsBetween(db *gorm.DB) *gorm.DB {
	return db.Where("dog_id > 50 and dog_id < ?", 100)
}

func GetBetween(c *fiber.Ctx) error {
	db := database.DBConn
	var dog []m.Dogs

	db.Scopes(DogsBetween).Find(&dog)

	return c.Status(200).JSON(dog)
}

func GetDogsJsonV2(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //10ตัว

	var dataResults []m.DogsRes
	sum_red := 0
	sum_green := 0
	sum_pink := 0
	sum_nocolor := 0
	for _, v := range dogs { //1 inet 112 //2 inet1 113
		typeStr := ""
		if v.DogID >= 10 && v.DogID <= 50 {
			typeStr = "red"
			sum_red++

		} else if v.DogID >= 100 && v.DogID <= 150 {
			typeStr = "green"
			sum_green++
		} else if v.DogID >= 200 && v.DogID <= 250 {
			typeStr = "pink"
			sum_pink++
		} else {
			typeStr = "no color"
			sum_nocolor++
		}

		d := m.DogsRes{
			Name:  v.Name,  //inet1
			DogID: v.DogID, //113
			Type:  typeStr, //green
		}
		dataResults = append(dataResults, d)
		// sumAmount += v.Amount
	}

	r := m.ResultDataV2{
		Data:    dataResults,
		Name:    "golang-test",
		Count:   len(dogs), //หาผลรวม,
		Red:     sum_red,
		Green:   sum_green,
		Pink:    sum_pink,
		Nocolor: sum_nocolor,
	}
	return c.Status(200).JSON(r)
}

// --------------------------------------------------------
func GetCompanies(c *fiber.Ctx) error {
	db := database.DBConn
	var company []m.Company

	//db.Scopes(DogIDGreaterThan100).Find(&company)
	db.Find(&company) //delelete = null
	return c.Status(200).JSON(company)
}

func GetCompany(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("email"))
	var cpn []m.Company

	result := db.Find(&cpn, "email = ?", search)

	// returns found records count, equals `len(users)
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&cpn)
}

func AddCompany(c *fiber.Ctx) error {
	//twst3
	db := database.DBConn
	var company m.Company

	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&company)
	return c.Status(201).JSON(company)
}

func UpdateCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Company
	id := c.Params("id")

	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&company)
	return c.Status(200).JSON(company)
}

func RemoveCompany(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var cpn m.Company

	result := db.Delete(&cpn, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

//----------------------------------------------------------------------

// MINI PROJECT
