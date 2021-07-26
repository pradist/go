package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n=============================================\n", sql)
}

var db *gorm.DB

func main() {
	dsn := "root:P@ssw0rd@tcp(localhost:3306)/pooh?parseTime=True"
	dial := mysql.Open(dsn)
	var err error
	db, err = gorm.Open(dial, &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: false,
	})
	if err != nil {
		panic(err)
	}
	// CreateGender("Female")
	// db.AutoMigrate(Gender{}, Test{}, Customer{})
	// GetGenders()
	// GetGenderByName("Male")
	// Update2(1, "Male1")
	// DeleteGender(3)
	// CreateMyTest(0, "test1")
	// CreateMyTest(0, "test2")
	// CreateMyTest(0, "test3")

	// DeleteTest(3)

	// db.Migrator().CreateTable(Customer{})

	GetCustomers()
}

func CreateCustomer(name string, genderID uint) {
	customer := Customer{
		Name:     name,
		GenderID: genderID,
	}
	tx := db.Create(&customer)
	if tx.Error != nil {
		fmt.Printf("%v\n", tx.Error)
		return
	}
	fmt.Println(customer)
}

func GetCustomers() {
	customers := []Customer{}
	tx := db.Preload("Gender").Find(&customers)
	if tx.Error != nil {
		fmt.Printf("%v\n", tx.Error)
		return
	}
	fmt.Println(customers)
}

func CreateMyTest(code uint, name string) {
	test := Test{
		Code: code,
		Name: name,
	}
	tx := db.Create(&test)
	if tx.Error != nil {
		fmt.Printf("%v\n", tx.Error)
		return
	}
	fmt.Print(test)
}

func GetTests() {
	tests := []Test{}
	tx := db.Find(&tests)
	if tx.Error != nil {
		fmt.Printf("%v\n", tx.Error)
		return
	}
	for _, t := range tests {
		fmt.Println(t.Code, t.Name)
	}
}

func DeleteTest(id uint) {
	tx := db.Delete(&Test{}, id)
	if tx.Error != nil {
		fmt.Printf("%v\n", tx.Error)
		return
	}
}

func DeleteGender(id uint) {
	tx := db.Delete(&Gender{}, id)
	if tx.Error != nil {
		fmt.Printf("%v\n", tx.Error)
		return
	}
	fmt.Println("Deleted")
	GetGender(3)
}

func Update2(id uint, name string) {
	gender := Gender{Name: name}
	tx := db.Model(&Gender{}).Where("id", id).Updates(gender)
	if tx.Error != nil {
		fmt.Printf("%v\n", tx.Error)
		return
	}
	GetGender(id)
}

func Update(id uint, name string) {
	gender := Gender{}
	tx := db.First(&gender, id)
	if tx.Error != nil {
		fmt.Printf("%v\n", tx.Error)
		return
	}
	gender.Name = name
	tx = db.Save(&gender)
	if tx.Error != nil {
		fmt.Printf("%v\n", tx.Error)
		return
	}
	GetGender(id)
}

func GetGenders() {
	genders := []Gender{}
	tx := db.Order("id").Find(&genders)
	if tx.Error != nil {
		fmt.Printf("%v\n", tx.Error)
		return
	}
	fmt.Print(genders)
}

func GetGender(id uint) {
	gender := Gender{}
	tx := db.First(&gender, id)
	if tx.Error != nil {
		fmt.Printf("%v\n", tx.Error)
		return
	}
	fmt.Print(gender)
}

func GetGenderByName(name string) {
	gender := Gender{}
	// tx := db.First(&gender, "name", name)
	tx := db.Where("name", name).First(&gender)
	if tx.Error != nil {
		fmt.Printf("%v\n", tx.Error)
		return
	}
	fmt.Print(gender)
}

func CreateGender(name string) {
	gender := Gender{
		Name: name,
	}
	tx := db.Create(&gender)
	if tx.Error != nil {
		fmt.Printf("%v\n", tx.Error)
		return
	}
	fmt.Print(gender)
}

type Customer struct {
	ID       uint
	Name     string
	Gender   Gender
	GenderID uint
}

type Gender struct {
	ID   uint
	Name string `gorm:"unique;size:10"`
}

type Test struct {
	gorm.Model
	Code uint
	Name string `gorm:"column:myname; type:varchar(20); unique; default:Hello; not null"`
}

func (t Test) TableName() string {
	return "MyTest"
}
