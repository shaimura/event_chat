package db

import (
	"fmt"
	"os"
	"time"

	"example.com/go-mod/app/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB

var err error

func Connect() {

	// envファイルを読み込む
	enverr := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV")))
	if enverr != nil {
		fmt.Println("ファイルの読み込みに失敗しました")
	}

	// 読み込んだenvファイルから値を取得する
	DBMS := os.Getenv("DBMS")
	USER := os.Getenv("USER")
	PASS := os.Getenv("PASS")
	DBNAME := os.Getenv("DBNAME")
	PROTOCOL := os.Getenv("PROTOCOL")

	// mysqlの設定を変数に代入する
	// dockerの場合、user:password@tcp(container-name:port)/dbname ※mysql はデフォルトで用意されているDB
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	// データベースに接続する
	db, err = gorm.Open(DBMS, CONNECT)
	// データベースに接続できなかった場合の処理
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("データベースに接続しました")
	}

	db.LogMode(true)

	// コネクションプールの接続を制限する
	db.DB().SetConnMaxLifetime(time.Minute * 3) // ドライバーによって接続が安全に閉じられるようにするために必要
	db.DB().SetMaxOpenConns(10)                 // アプリケーションが使用する接続の数を制限する
	db.DB().SetMaxIdleConns(10)                 // SetMaxIdleConnsはアイドル状態のコネクションプール内の最大数を設定する

	db.DropTable(&model.User{}, &model.Accesstoken{}, &model.Userchatroom{}, &model.UserMessage{})
	db.AutoMigrate(&model.User{}, &model.Accesstoken{}, &model.Userchatroom{}, &model.UserMessage{})

	// var allmodel = &model.User{}, &model.Accesstoken{}
	// db.DropTable(allmodel)
	// db.AutoMigrate(allmodel)

}

// データベースの情報を取得する
func Get() *gorm.DB {
	return db
}

// データベースの接続を切断する
func Close() {
	db.Close()
}
