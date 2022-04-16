package storage

import (
	"UserAuth/common"
	"UserAuth/models"
	"time"

	"fmt"

	"github.com/dgrijalva/jwt-go"
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

func (c *cursor) VerifyUser(cs *models.User) map[string]interface{} {

	user := new(models.User)
	fmt.Println("Verify user cs.UserName:", cs.UserName)
	fmt.Println("Verify user variable user:", user.UserName)
	//verify username:
	if c.Db.Where("user_name = ?", cs.UserName).First(user).RecordNotFound() {
		return map[string]interface{}{"message": "User not found"}
	}

	fmt.Println("after where, user.username:", user.UserName)
	fmt.Println("after where, user.password:", user.Password)
	//verify Password
	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cs.Password))
	if passErr != nil && passErr == bcrypt.ErrMismatchedHashAndPassword {
		//Password does not match!
		return map[string]interface{}{"message": "Wrong Password"}
	}

	//sign token
	tokenContent := jwt.MapClaims{
		"user":   cs.UserName,
		"expiry": time.Now().Add(time.Minute * 100).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	if err != nil {
		fmt.Println("Error in token generation: ", err)
		panic(err)
	}

	var resp = map[string]interface{}{"message": "Successfully Logged In"}
	resp["jwt"] = token

	return resp
}
