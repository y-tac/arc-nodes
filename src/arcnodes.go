import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "./action"
	"./policy"
)
func get() {

}
func main() {
    // init conf file
    // c

    // Echoのインスタンス作る
    e := echo.New()

    // 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // ルーティング
    e.GET("/hello", handler.MainPage())

    // サーバー起動
    e.Start(":1323")    //ポート番号指定してね
}