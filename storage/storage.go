package storage

import (
	"UserAuth/common"
	"UserAuth/models"

	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type cursor struct {
	Db *gorm.DB
}

func GetCursor() *cursor {
	c := new(cursor)
	c.Db = common.Db
	return c
}

func (c *cursor) CreateUser(cs *models.User) (*models.User, error) {

	fmt.Println("Function CreateUser:")
	fmt.Println("username: ", cs.UserName)
	fmt.Println("password:", cs.Password)

	pass, err := bcrypt.GenerateFromPassword([]byte(cs.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println(err)
	}

	cs.Password = string(pass)

	fmt.Println("password After Hash:", cs.Password)

	result := c.Db.Create(cs)
	if err := result.Error; err != nil {
		return nil, err
	}

	record := result.Value.(*models.User)

	return record, nil
}

func (c *cursor) VerifyUser(cs *models.User) error {
	user := new(models.User)
	fmt.Println("Verify user cs.UserName:", cs.UserName)
	fmt.Println("Verify user variable user:", user.UserName)

	if err := c.Db.Where("user_name = ?", cs.UserName).First(user).Error; err != nil {
		return err
	}
	fmt.Println("after where, user.username:", user.UserName)
	fmt.Println("after where, user.password:", user.Password)

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cs.Password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		//var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return errf
	}
	return errf
}
