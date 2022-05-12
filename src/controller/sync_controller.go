package controller

import (
	"fmt"
	"password_manager/src/common/constant"
	"password_manager/src/common/response"
	"password_manager/src/dao"
	"password_manager/src/handler"
	"password_manager/src/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SyncController struct {
}

func SyncRegister(userGrp *gin.RouterGroup) {

	SyncController := &SyncController{}

	userGrp.Use().POST("/open", SyncController.open)

}

func (c SyncController) open(ctx *gin.Context) {

	user := ctx.Keys["user"].(*handler.UserClaims)

	status, _ := ctx.GetPostForm("status")
	num, _ := strconv.Atoi(status)

	if num == 0 {

		dao.UpdateDb(false)

		utils.DeleteJob(user.ID)
	} else {

		dao.UpdateDb(true)

		utils.AddJob(sync, "*/1 * * * *", user.ID)
	}

	response.Success(ctx, constant.SelectSuccessCode, constant.SelectSuccessMsg, nil)
}

func sync() {

	fmt.Println("数据同步功能定时任务 run。。。。。")
}
