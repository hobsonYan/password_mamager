package route

import (
	"password_manager/src/controller"

	"github.com/gin-gonic/gin"
)

func PathRoute(r *gin.Engine) *gin.Engine {

	rootPath := r.Group("/gin")

	{

		userPath := rootPath.Group("/user")

		{
			controller.UserRegister(userPath)
		}

		passwordPath := rootPath.Group("/passward")

		{
			controller.PasswordRegister(passwordPath)
		}

	}

	return r
}
