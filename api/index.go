package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"health/models"
	"strings"
)

func Version(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"status":       1,
			"version":      111,
			"upgrade_path": "http://www.test.com",
			"note":         "1.更新啦",
		},
		"msg": "this is message",
	})
	return
}
func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, gin.H{
		"status":  1,
		"message": "Success",
		"data":    data,
	})
}
func Fail(ctx *gin.Context, msg string) {
	ctx.JSON(200, gin.H{
		"status":  0,
		"message": msg,
		"data":    "",
	})
}

func processErr(m interface{}, err error) string {
	if err == nil {
		return ""
	}
	invalid, ok := err.(*validator.InvalidValidationError)
	if ok {
		return "参数错误：" + invalid.Error()
	}
	unmarshal, ok := err.(*json.UnmarshalTypeError)
	if ok {
		return "参数类型错误：" + unmarshal.Field
	}

	validErr := err.(validator.ValidationErrors)
	for _, info := range validErr {
		var msg string
		switch info.ActualTag() {
		case "required":
			msg = "不能为空"
			break
		case "phone":
			msg = "手机号格式错误"
			break
		case "min":
			msg = "长度不能小于 " + info.Param() + " 个字符"
			break
		case "max":
			msg = "长度不能大于 " + info.Param() + " 个字符"
		}
		return info.Field() + " : " + msg

	}
	return ""
}

// Login 用户注册
func Login(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBind(&user); err != nil {
		Fail(ctx, processErr(user, err))
		return
	}
	u, err := models.Login(user)
	if err != nil {
		Fail(ctx, err.Error())
		return
	}
	token, err := models.GenToken(*u)
	if err != nil {
		Fail(ctx, err.Error())
		return
	}
	Success(ctx, gin.H{"token": token})
}

// Register 用户注册
func Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBind(&user); err != nil {
		Fail(ctx, processErr(user, err))
		return
	}
	err := models.Register(user)
	if err != nil {
		Fail(ctx, err.Error())
		return
	}
	Success(ctx, nil)
}

// List 数据列表
func List(ctx *gin.Context) {
	var cond models.Cond
	ctx.ShouldBind(&cond)
	if v, ok := ctx.Get("uid"); ok {
		cond.Uid = v.(uint32)
	}
	ret, err := models.List(cond)
	if err != nil {
		Fail(ctx, err.Error())
		return
	}
	var data []interface{}
	for _, item := range ret {
		day, _, _ := strings.Cut(item.Day, "T")
		data = append(data, gin.H{
			"id":     item.ID,
			"uid":    item.Uid,
			"weight": item.Weight,
			"waist":  item.Waist,
			"bust":   item.Bust,
			"hip":    item.Hip,
			"thigh":  item.Thigh,
			"arm":    item.Arm,
			"day":    day,
		})
	}
	Success(ctx, data)
}

// Create 添加数据
func Create(ctx *gin.Context) {
	var data models.Health
	if err := ctx.ShouldBind(&data); err != nil {
		Fail(ctx, processErr(data, err))
		return
	}
	if v, ok := ctx.Get("uid"); ok {
		data.Uid = v.(uint32)
	}
	ret, err := models.Create(data)
	if err != nil {
		Fail(ctx, err.Error())
		return
	}
	Success(ctx, ret)

}

// Del 删除数据
func Del(ctx *gin.Context) {
	var cond models.Cond
	ctx.ShouldBind(&cond)
	if cond.Id == 0 {
		Fail(ctx, "参数有误")
		return
	}
	if v, ok := ctx.Get("uid"); ok {
		cond.Uid = v.(uint32)
	}
	if err := models.Del(cond); err != nil {
		Fail(ctx, err.Error())
		return
	}
	Success(ctx, nil)
}
