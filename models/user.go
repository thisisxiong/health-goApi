package models

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"health/global"
	"regexp"
	"time"
)

var tokenSignKey = "daydayup"

type User struct {
	Base
	Name     string `form:"name" json:"name"`
	Phone    string `gorm:"unique" form:"phone" json:"phone" binding:"required,phone"`
	Password string `form:"password" json:"password" binding:"required,min=6"`
}

var checkphone validator.Func = func(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	if match, _ := regexp.Match(`^1[3-9]\d{9}$`, []byte(phone)); match {
		return true
	}
	return false
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("phone", checkphone)
	}
}

func Login(user User) (*User, error) {
	var u User
	if result := global.Db.Where(&User{Phone: user.Phone}).First(&u); result.RowsAffected == 0 {
		return nil, errors.New("用户不存在")
	}
	if u.Password != user.Password {
		return nil, errors.New("密码错误")
	}
	return &u, nil
}

func Register(user User) error {
	var u User
	if result := global.Db.Where(&User{Phone: user.Phone}).First(&u); result.RowsAffected > 0 {
		return errors.New("用户已存在")
	}
	if user.Name == "" {
		user.Name = user.Phone
	}
	if result := global.Db.Create(&user); result.RowsAffected == 0 {
		return errors.New("注册失败，请重试")
	}
	return nil
}

type CustomClaims struct {
	Id    int
	Phone string
	*jwt.RegisteredClaims
}

func GenToken(user User) (string, error) {
	claims := &CustomClaims{
		Id:    int(user.ID),
		Phone: user.Phone,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "health",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(tokenSignKey))
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSignKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token invalid")
}
