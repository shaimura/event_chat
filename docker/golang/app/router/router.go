package router

import (
	"example.com/go-mod/app/controller"
	"example.com/go-mod/app/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() {

	router := gin.Default()

	// CORS対応
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	router.Use(cors.New(config))

	userhandler := controller.UserHandler{
		db.Get(),
	}

	chathandler := controller.ChatHandler{
		db.Get(),
	}

	router.GET("/test", controller.Test)
	router.POST("/signup", userhandler.SignUp)                    // ユーザー登録
	router.POST("/signin", userhandler.SignIn)                    // サインイン
	router.POST("/gettoken", userhandler.GetToken)                // トークン取得
	router.POST("/refreshidtoken", userhandler.RefreshIdToken)    // トークン更新
	router.POST("/userchatroom", userhandler.UserChatRoom)        // トークン更新
	router.GET("/getuser/:id", userhandler.GetUser)               // ユーザー取得
	router.GET("/allusers", userhandler.ALLUsers)                 // ユーザー取得
	router.POST("/getusers", userhandler.GetUsers)                // ユーザー取得
	router.POST("/sendusermessage", chathandler.SendUserMessage)  // ユーザー取得
	router.GET("/getusermessage/:id", chathandler.GetUserMessage) // ユーザー取得
	router.GET("/ws", controller.Chat)
	go controller.HandleMessages()

	router.Run(":8888")

}
