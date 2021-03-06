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
	userGrp.Use().POST("/update", UserController.update)

	userGrp.Use().POST("/netLogin", UserController.netLogin)
}

func (c UserController) findUser(ctx *gin.Context) {
	username := ctx.Query("username")
	userString := dao.View("user", username)
	user := model.User{}
	json.Unmarshal([]byte(userString), &user)
	userVo := model.UserVo{Id: user.Id, Username: user.Username, Phone: user.Phone}
	response.Success(ctx, constant.SelectSuccessCode, constant.SelectSuccessMsg, userVo)
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
					Name:           user.Username,
					Phone:          user.Phone,
					StandardClaims: jwt.StandardClaims{},
				}),
				"id":       user.Id,
				"username": user.Username,
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

		userString := dao.View("user", form.Username)

		if len(userString) > 0 {
			response.Failure(ctx, constant.SelectFailureCode, "用户名已被占用", nil)
		} else {
			id, _ := uuid.NewV4()
			user, _ := json.Marshal(model.User{Id: id.String(), Username: form.Username, Password: form.Password, Phone: form.Phone, Time: time.Now().UnixNano()})

			dao.Update("user", form.Username, string(user))

			response.Success(ctx, constant.SelectSuccessCode, "用户注册成功", nil)
		}
	}
}

func (c UserController) update(ctx *gin.Context) {
	var form model.UpdateForm

	if err := ctx.ShouldBind(&form); err != nil {
		response.Failure(ctx, constant.SelectFailureCode, "入参绑定失败", nil)
	} else {
		user := ctx.Keys["user"].(*handler.UserClaims)

		userString := dao.View("user", user.Name)
		u := model.User{}
		json.Unmarshal([]byte(userString), &u)

		if len(form.New_Password) > 0 && form.New_Password != u.Password {
			u.Password = form.New_Password
		}

		if len(form.New_Phone) > 0 && form.New_Phone != u.Phone {
			u.Phone = form.New_Phone
		}

		if len(form.New_Username) > 0 && form.New_Username != u.Username {
			u.Username = form.New_Username

			dao.Delete("user", user.Name)
		}

		newUser, _ := json.Marshal(model.User{Id: u.Id, Username: u.Username, Password: u.Password, Phone: u.Phone, Time: time.Now().UnixNano()})
		dao.Update("user", u.Username, string(newUser))

		go syncUser(&u)

		response.Success(ctx, constant.SelectSuccessCode, "用户更新成功", nil)
	}
}

func (c UserController) netLogin(ctx *gin.Context) {
	var form model.LoginForm

	if err := ctx.ShouldBind(&form); err != nil {
		response.Failure(ctx, constant.SelectFailureCode, "入参绑定失败", nil)
	} else {
		db := dao.GetDb()
		var user model.User
		db.Where("username = ? and password = ?", form.Username, form.Password).Debug().Find(&user)

		if user == (model.User{}) {
			response.Failure(ctx, constant.SelectFailureCode, "用户名或密码不正确", nil)
		} else {

			syncUser(&user)

			syncPassword(user.Id)

			response.Success(ctx, constant.SelectSuccessCode, constant.SelectSuccessMsg, nil)
		}
	}
}
