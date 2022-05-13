package controller

import (
	"encoding/json"
	"password_manager/src/common/constant"
	"password_manager/src/common/response"
	"password_manager/src/dao"
	"password_manager/src/handler"
	"password_manager/src/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dianjiu/gokit/uuid"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func UserRegister(userGrp *gin.RouterGroup) {

	UserController := &UserController{}
	userGrp.Use().GET("/findUser", UserController.findUser)
	userGrp.Use().POST("/login", UserController.login)
	userGrp.Use().POST("/register", UserController.register)
}

func (c UserController) findUser(ctx *gin.Context) {
	username := ctx.Query("username")
	userString := dao.View("user", username)
	user := model.User{}
	json.Unmarshal([]byte(userString), &user)
	response.Success(ctx, constant.SelectSuccessCode, constant.SelectSuccessMsg, user)
}

func (c UserController) login(ctx *gin.Context) {
	var form model.LoginForm

	if err := ctx.ShouldBind(&form); err != nil {
		response.Failure(ctx, constant.SelectFailureCode, "入参绑定失败", nil)
	} else {
		userString := dao.View("user", form.Username)
		user := model.User{}
		json.Unmarshal([]byte(userString), &user)
		if user.Password != form.Password {
			response.Failure(ctx, constant.SelectFailureCode, "密码错误", nil)
		} else {
			response.Success(ctx, constant.SelectSuccessCode, constant.SelectSuccessMsg, gin.H{
				"token": handler.GenerateToken(&handler.UserClaims{
					ID:             user.Id,
					Name:           user.Usrename,
					Phone:          user.Phone,
					StandardClaims: jwt.StandardClaims{},
				}),
				"id":       user.Id,
				"username": user.Usrename,
				"phone":    user.Phone,
			})
		}
	}
}

func (c UserController) register(ctx *gin.Context) {
	var form model.RegisterForm
	if err := ctx.ShouldBind(&form); err != nil {
		response.Failure(ctx, constant.SelectFailureCode, "入参绑定失败", nil)
	} else {

		userString := dao.View("user", form.Usrename)

		if len(userString) > 0 {
			response.Failure(ctx, constant.SelectFailureCode, "用户名已被占用", nil)
		} else {
			id, _ := uuid.NewV4()
			user, _ := json.Marshal(model.User{Id: id.String(), Usrename: form.Usrename, Password: form.Password, Phone: form.Phone, Time: time.Now().UnixNano()})

			dao.Update("user", form.Usrename, string(user))

			response.Success(ctx, constant.SelectSuccessCode, "用户注册成功", nil)
		}
	}
}
