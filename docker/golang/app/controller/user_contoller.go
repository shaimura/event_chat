package controller

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"

	"example.com/go-mod/app/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)

type UserHandler struct {
	Db *gorm.DB
}

var err error

func SignUpError(user model.User, password string) []string {
	// バリデーションの処理
	// バリデーションの準備
	var validate *validator.Validate

	validate = validator.New()
	var errormessage []string
	var errorfield string
	var errortag string

	if password == "" {
		errormessage = append(errormessage, "パスワードを入力してください\n")
	} else if len(password) < 4 || len(password) > 8 {
		errormessage = append(errormessage, "パスワードは4〜8文字で入力してください\n")
	}

	errs := validate.Struct(user)
	if errs != nil {
		if _, ok := errs.(*validator.ValidationErrors); ok {
			return errormessage
		}

		for _, err := range errs.(validator.ValidationErrors) {
			errorfield = err.Field()
			var setfield string
			if errorfield == "Username" {
				setfield = "ユーザー名"
			} else if errorfield == "Email" {
				setfield = "メールアドレス"
			}
			if errortag == "required" {
				errormessage = append(errormessage, (setfield + "を入力してください\n"))
			}
		}
		return errormessage
	}
	return errormessage
}

func (handler *UserHandler) SignUp(c *gin.Context) {
	accesstoken := model.Accesstoken{}

	user := model.User{}
	username, _ := c.GetPostForm("username")
	username = strings.TrimSpace(username)

	password, _ := c.GetPostForm("password")
	password = strings.TrimSpace(password)

	email, _ := c.GetPostForm("email")
	email = strings.TrimSpace(email)

	user.Username = username
	user.Email = email

	// エラー判定
	errormessage := SignUpError(user, password)
	if errormessage != nil {
		c.JSON(200, gin.H{
			"errormessage": errormessage,
			"user":         user,
			"accesstoken":  accesstoken,
		})
		return
	}

	// パスワードをハッシュ値に変更
	setpassword := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
	user.Password = setpassword

	// ユーザートークンを作成する
	seedtoken := username + password + email
	usertoken := fmt.Sprintf("%x", seedtoken)
	user.Usertoken = usertoken

	if errormessage == nil {
		// 入力情報が登録されているか確認する
		if err := handler.Db.Create(&user).Error; err != nil {
			errormessage = append(errormessage, "入力された情報はすでに登録されています")
			c.JSON(200, gin.H{
				"errormessage": errormessage,
				"user":         user,
				"accesstoken":  accesstoken,
			})
			return
		}
	}

	// アクセストークンを取得する
	// accesstoken := model.Accesstoken{}
	accessuser := model.User{}
	id := user.ID
	handler.Db.First(&accessuser, id)
	now := time.Now()
	hour := now.Add(1 * time.Hour)
	nowUTC := hour.UTC()

	expirationdate := (nowUTC.UnixNano() / int64(time.Millisecond)) // ミリ秒を算出する

	accessseedtoken := usertoken + fmt.Sprint(expirationdate)
	accesssettoken := fmt.Sprintf("%x", sha256.Sum256([]byte(accessseedtoken)))

	accesstoken.UserID = user.ID
	accesstoken.Username = user.Username
	accesstoken.Accesstoken = accesssettoken
	accesstoken.Expirationdata = expirationdate

	handler.Db.Create(&accesstoken)

	c.JSON(200, gin.H{
		"errormessage": errormessage,
		"user":         user,
		"accesstoken":  accesstoken,
	})
}

// ログイン機能
func (handler *UserHandler) SignIn(c *gin.Context) {
	user := model.User{}
	var message string

	username, _ := c.GetPostForm("username")
	email, _ := c.GetPostForm("email")
	password, _ := c.GetPostForm("password")

	searchtoken := fmt.Sprintf("%x", (username + password + email))

	handler.Db.Where("Usertoken = ?", searchtoken).Find(&user)

	if user.Usertoken == searchtoken {
		message = "ログインしました"
	}

	c.JSON(200, gin.H{
		"message": message,
		"user":    user,
	})
}

// アクセストークン更新
func (handler *UserHandler) RefreshIdToken(c *gin.Context) {
	accesstoken := model.Accesstoken{}

	userid, _ := c.GetPostForm("userid")

	fmt.Println(userid)

	handler.Db.Where("user_id = ?", userid).Find(&accesstoken)

	settoken := accesstoken.Accesstoken

	now := time.Now()
	hour := now.Add(1 * time.Hour)
	nowUTC := hour.UTC()

	expirationdate := (nowUTC.UnixNano() / int64(time.Millisecond))

	seedtoken := settoken + fmt.Sprint(expirationdate)
	newtoken := fmt.Sprintf("%x", sha256.Sum256([]byte(seedtoken)))

	accesstoken.Accesstoken = newtoken
	accesstoken.Expirationdata = expirationdate

	handler.Db.Model(&accesstoken).Updates(&accesstoken)

	c.JSON(200, accesstoken)
}

// ユーザー情報取得
func (handler *UserHandler) GetUser(c *gin.Context) {
	user := model.User{}
	userid := c.Param("id")

	handler.Db.First(&user, userid)
	c.JSON(200, user)
}

// 使っていない
func (handler *UserHandler) GetToken(c *gin.Context) {
	accesstoken := model.Accesstoken{}
	user := model.User{}

	fmt.Println("start")

	id, _ := c.GetPostForm("id")
	handler.Db.First(&user, id)

	usertoken := user.Usertoken

	now := time.Now()
	hour := now.Add(1 * time.Hour)
	nowUTC := hour.UTC()

	expirationdate := (nowUTC.UnixNano() / int64(time.Millisecond)) // ミリ秒を算出する

	seedtoken := usertoken + fmt.Sprint(expirationdate)
	settoken := fmt.Sprintf("%x", sha256.Sum256([]byte(seedtoken)))

	accesstoken.UserID = user.ID
	accesstoken.Username = user.Username
	accesstoken.Accesstoken = settoken
	accesstoken.Expirationdata = expirationdate

	handler.Db.Create(&accesstoken)

	fmt.Println(accesstoken)
	c.JSON(200, accesstoken)
}

// 全ユーザー情報取得
func (handler *UserHandler) ALLUsers(c *gin.Context) {
	var users []model.User
	handler.Db.Find(&users)
	c.JSON(200, users)
}

func (handler *UserHandler) UserChatRoom(c *gin.Context) {

	userchatroom := model.Userchatroom{}
	curentuserid, _ := c.GetPostForm("curentuserid")
	seconduserid, _ := c.GetPostForm("seconduserid")

	intcurentuserid, _ := strconv.Atoi(curentuserid)
	setcurentuserid := uint(intcurentuserid)

	intseconduserid, _ := strconv.Atoi(seconduserid)
	setseconduserid := uint(intseconduserid)

	if err := handler.Db.Where("Firstuserid in (?) AND Seconduserid in (?)", []uint{setcurentuserid, setseconduserid}, []uint{setcurentuserid, setseconduserid}).Find(&userchatroom).Error; err != nil {
		fmt.Println("err")
		userchatroom.Firstuserid = setcurentuserid
		userchatroom.Seconduserid = setseconduserid
		handler.Db.Create(&userchatroom)
	}

	c.JSON(200, userchatroom)
}

func (handler *UserHandler) GetUsers(c *gin.Context) {
	var users []model.User
	userid, _ := c.GetPostForm("id")
	fmt.Println(userid)

	handler.Db.Where("ID <> ?", userid).Find(&users)

	c.JSON(200, users)
}
