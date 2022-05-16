package controller

import (
	"encoding/json"
	"log"
	"password_manager/src/common/constant"
	"password_manager/src/common/response"
	"password_manager/src/dao"
	"password_manager/src/handler"
	"password_manager/src/model"
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

		utils.AddJob(syncPassword, "*/1 * * * *", user.ID)
	}

	response.Success(ctx, constant.SelectSuccessCode, constant.SelectSuccessMsg, nil)
}

func syncUser(user *model.User) {

	db := dao.GetDb()

	// 人员信息同步
	userString := dao.View("user", user.Username)
	boltUser := model.User{}
	json.Unmarshal([]byte(userString), &boltUser)

	var mysqlUser model.User
	db.Where("id = ?", user.Id).Debug().Find(&mysqlUser)

	if (boltUser == (model.User{}) && mysqlUser != (model.User{})) || mysqlUser.Time > boltUser.Time {

		u, _ := json.Marshal(model.User{Id: mysqlUser.Id, Username: mysqlUser.Username, Password: mysqlUser.Password, Phone: mysqlUser.Phone, Time: mysqlUser.Time})

		dao.Update("user", mysqlUser.Username, string(u))
	}

	if boltUser != (model.User{}) && mysqlUser == (model.User{}) {

		db.Create(boltUser)
	}

	if mysqlUser.Time < boltUser.Time {

		db.Model(model.User{}).Where("id = ?", boltUser.Id).Updates(&boltUser)
	}
}

func syncPassword(userId string) {

	log.Println("数据同步功能定时任务 run。。。。。")
	db := dao.GetDb()

	// 密码信息同步
	strList, total := dao.List(userId, "", 0, 0)
	deleteList, num := dao.List("delete_password", userId, 0, 0)

	boltPasswordList := make([]model.Password, 0)
	boltMap := make(map[string]model.Password)
	for i := 0; i < total; i++ {
		password := model.Password{}
		json.Unmarshal([]byte(strList[i]), &password)
		boltMap[password.ID] = password
		boltPasswordList = append(boltPasswordList, password)
	}

	var mysqlPasswordList []model.Password
	db.Where("user_id = ?", userId).Debug().Find(&mysqlPasswordList)
	mysqlMap := make(map[string]model.Password)
	for i := 0; i < len(mysqlPasswordList); i++ {
		mysqlMap[mysqlPasswordList[i].ID] = mysqlPasswordList[i]
	}

	boltUpdateList := make([]model.Password, 0)
	boltDeleteList := make([]model.Password, 0)

	for i := 0; i < len(mysqlPasswordList); i++ {
		temp, ok := boltMap[mysqlPasswordList[i].ID]

		if !ok && mysqlPasswordList[i].Is_Delete == 0 {
			boltUpdateList = append(boltUpdateList, mysqlPasswordList[i])
		}

		if ok && mysqlPasswordList[i].Is_Delete == 0 && temp.Time < mysqlPasswordList[i].Time {
			boltUpdateList = append(boltUpdateList, mysqlPasswordList[i])
		}

		if ok && mysqlPasswordList[i].Is_Delete == 1 && temp.Is_Delete == 0 {
			boltDeleteList = append(boltDeleteList, mysqlPasswordList[i])
		}
	}

	mysqlUpdateList := make([]model.Password, 0)
	mysqlAddList := make([]model.Password, 0)
	mysqlDeleteList := make([]model.Password, 0)

	for i := 0; i < len(boltPasswordList); i++ {
		temp, ok := mysqlMap[boltPasswordList[i].ID]

		if !ok {
			mysqlAddList = append(mysqlAddList, boltPasswordList[i])
		}

		if ok && temp.Time < boltPasswordList[i].Time {
			mysqlUpdateList = append(mysqlUpdateList, boltPasswordList[i])
		}
	}
	for i := 0; i < num; i++ {
		if temp, ok := mysqlMap[deleteList[i]]; ok {
			mysqlDeleteList = append(mysqlDeleteList, temp)
		}
	}

	// bolt db 数据更新
	if len(boltUpdateList) > 0 {
		for i := 0; i < len(boltUpdateList); i++ {
			password, _ := json.Marshal(model.Password{ID: boltUpdateList[i].ID, Key: boltUpdateList[i].Key, Value: boltUpdateList[i].Value, Time: boltUpdateList[i].Time})
			dao.Update(userId, boltUpdateList[i].ID, string(password))
		}
	}
	if len(boltDeleteList) > 0 {
		for i := 0; i < len(boltDeleteList); i++ {
			dao.Delete(userId, boltDeleteList[i].ID)
			dao.Update("delete_password", userId, boltDeleteList[i].ID)
		}
	}

	// mysql 数据更新
	if len(mysqlAddList) > 0 {
		for i := 0; i < len(mysqlAddList); i++ {
			mysqlAddList[i].User_ID = userId
			db.Create(mysqlAddList[i])
		}
	}
	if len(mysqlUpdateList) > 0 {
		for i := 0; i < len(mysqlUpdateList); i++ {
			db.Model(model.Password{}).Where("id = ?", mysqlUpdateList[i].ID).Updates(&mysqlUpdateList[i])
		}
	}
	if len(mysqlDeleteList) > 0 {
		for i := 0; i < len(mysqlDeleteList); i++ {
			mysqlDeleteList[i].Is_Delete = 1
			db.Model(model.Password{}).Where("id = ?", mysqlDeleteList[i].ID).Updates(mysqlDeleteList[i])
		}
	}
}
