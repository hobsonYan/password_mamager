package controller

import (
	"encoding/json"
	"math/rand"
	"password_manager/src/common/constant"
	"password_manager/src/common/response"
	"password_manager/src/dao"
	"password_manager/src/handler"
	"password_manager/src/model"
	"strconv"
	"time"

	"github.com/dianjiu/gokit/uuid"
	"github.com/gin-gonic/gin"
)

type PasswordController struct {
}

func PasswordRegister(userGrp *gin.RouterGroup) {

	PasswordController := &PasswordController{}

	userGrp.Use().POST("/add", PasswordController.add)
	userGrp.Use().GET("/get", PasswordController.get)
	userGrp.Use().POST("/delete", PasswordController.delete)
	userGrp.Use().GET("/list", PasswordController.list)
	userGrp.Use().GET("/generate", PasswordController.generate)
}

func (c PasswordController) add(ctx *gin.Context) {

	user := ctx.Keys["user"].(*handler.UserClaims)

	var form model.AddForm
	if err := ctx.ShouldBind(&form); err != nil {
		response.Failure(ctx, constant.SelectFailureCode, "入参绑定失败", nil)
	} else {

		time := time.Now().UnixNano()

		if len(form.Id) > 0 {
			// 更新操作
			password, _ := json.Marshal(model.Password{ID: form.Id, Key: form.Key, Value: form.Value, Time: time})
			dao.Update(user.ID, form.Id, string(password))
			response.Success(ctx, constant.SelectSuccessCode, "密码修改成功", nil)
		} else {
			// 新增操作
			id, _ := uuid.NewV4()
			password, _ := json.Marshal(model.Password{ID: id.String(), Key: form.Key, Value: form.Value, Time: time})
			dao.Update(user.ID, id.String(), string(password))
			response.Success(ctx, constant.SelectSuccessCode, "密码保存成功", nil)
		}
	}

}

func (c PasswordController) get(ctx *gin.Context) {
	user := ctx.Keys["user"].(*handler.UserClaims)

	id := ctx.Query("id")
	passwordStr := dao.View(user.ID, id)
	password := model.Password{}
	json.Unmarshal([]byte(passwordStr), &password)
	response.Success(ctx, constant.SelectSuccessCode, constant.SelectSuccessMsg, password.Value)
}

func (c PasswordController) delete(ctx *gin.Context) {
	user := ctx.Keys["user"].(*handler.UserClaims)

	id := ctx.PostForm("id")
	dao.Delete(user.ID, id)

	dao.Update("delete_password", user.ID, id)

	response.Success(ctx, constant.SelectSuccessCode, "密码删除成功", nil)
}

func (c PasswordController) list(ctx *gin.Context) {
	user := ctx.Keys["user"].(*handler.UserClaims)

	var form model.ListForm
	if err := ctx.ShouldBind(&form); err != nil {
		response.Failure(ctx, constant.SelectFailureCode, "入参绑定失败", nil)
	} else {
		list, total := dao.List(user.ID, "", form.PageNum, form.PageSize)

		var voList []model.ListVo
		for i := 0; i < len(list); i++ {
			password := model.Password{}
			json.Unmarshal([]byte(list[i]), &password)
			vo := model.ListVo{ID: password.ID, Key: password.Key}
			voList = append(voList, vo)
		}

		response.Success(ctx, constant.SelectSuccessCode, constant.SelectSuccessMsg, gin.H{
			"total": total,
			"list":  voList,
		})
	}
}

func (c PasswordController) generate(ctx *gin.Context) {

	var str string
	num, _ := strconv.Atoi(ctx.Query("num"))

	max := 126
	min := 33

	rand.Seed(time.Now().Unix())
	for i := 0; i < num; i++ {
		n := rand.Intn(max-min) + min
		char := string(byte(n))
		str = str + char
	}

	response.Success(ctx, constant.SelectSuccessCode, constant.SelectSuccessMsg, str)
}
