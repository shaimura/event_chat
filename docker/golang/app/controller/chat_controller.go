package controller

import (
	"log"
	"net/http"
	"strconv"

	"example.com/go-mod/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
)

type ChatHandler struct {
	Db *gorm.DB
}

// 接続されているユーザー
var users = make(map[*websocket.Conn]bool)

// メッセージのブロードキャストチャンネル
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// メッセージの構造体
type Message struct {
	Userchatroomid int    `json:"userchatroomid"`
	Username       string `json:"username"`
	Message        string `json:"message"`
}

// 受け取ったリクエストを処理する
func HandleConnections(w http.ResponseWriter, r *http.Request) {

	// 送られてきたGETリクエストをwebsocketにアップグレード
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 処理が終わったらwebsocketのコネクションを閉じる
	defer ws.Close()

	// ユーザーを新しく登録する
	users[ws] = true

	for {
		var msg Message

		// 新しいメッセージをJSONとして読み込み、Messageオブジェクトにマッピングする
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(users, ws)
			break
		}

		// 新しく受診されたメッセージをブロードキャストチャンネルに送る
		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		// ブロードキャストから次のメッセージを受け取る
		msg := <-broadcast

		// 現在接続しているユーザー全てにメッセージを送信する
		for user := range users {
			err := user.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				user.Close()
				delete(users, user)
			}
		}
	}
}

func Chat(c *gin.Context) {
	HandleConnections(c.Writer, c.Request)
}

type SetMessage struct {
	User    model.User
	Message model.UserMessage
}

func (handler *ChatHandler) GetUserMessage(c *gin.Context) {
	var usermessages []model.UserMessage
	var userid uint
	var setmessages []SetMessage
	user := model.User{}
	setmessage := SetMessage{}

	userchatroomid := c.Param("id")
	handler.Db.Where("userchatroom_id = ?", userchatroomid).Find(&usermessages)

	for _, usermessage := range usermessages {
		userid = usermessage.UserID
		user = model.User{}
		setmessage = SetMessage{}

		handler.Db.Where("ID = ?", userid).Find(&user)

		setmessage.Message = usermessage
		setmessage.User = user

		setmessages = append(setmessages, setmessage)
	}

	c.JSON(200, setmessages)
}

func (handler *ChatHandler) SendUserMessage(c *gin.Context) {
	usermessage := model.UserMessage{}

	userid, _ := c.GetPostForm("userid")
	intuserid, _ := strconv.Atoi(userid)
	var setuserid uint
	setuserid = uint(intuserid)

	userchatroomid, _ := c.GetPostForm("userchatroomid")
	intuserchatroomid, _ := strconv.Atoi(userchatroomid)
	var setuserchatroomid uint
	setuserchatroomid = uint(intuserchatroomid)

	message, _ := c.GetPostForm("message")

	usermessage.UserID = setuserid
	usermessage.UserchatroomID = setuserchatroomid
	usermessage.Message = message

	handler.Db.Create(&usermessage)

}
