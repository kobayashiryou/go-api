package main

import (
	"go-api/config"
	"go-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // muxかな？
	config.Cors(router)
	config.DbInit()
	router.GET("/", controller.TodoAll)
	router.POST("/memo", controller.CreateTodo)
	router.POST("/", controller.DeleteTodo)
	router.POST("/update", controller.UpdateTodo)
	router.Run() // 0.0.0.0:8080 でサーバーを立てます。
}

// := は変数の定義と型の定義をしてくれる、役割的にはconstと同じで変数の再代入ができない
// func getAll() []mode.Todoの[]mode.Todoは戻り値の型

/*
今の疑問
ハンドラとハンドラー関数の違いは？？
*/

/*
ここではマルチプレクサを生成しています。マルチプレクサとは、リクエストされたURLを調べて、所定のハンドラ（関数）にリダイレクトしてくれる。
RubyとかでいうRoutesかな
*/
// func main() {

// // ここでマルチプレクサを生成している。Muxはマルチプレクサの略かな。
// mux := http.NewServeMux()
// /*
// 	マルチプレクサはハンドラへのリダイレクトだけでなく、静的なファイル返送にも使えます。
// 	以下ではpublic配下にあるファイルを配信するハンドラを作成しています
// */
// files := http.FileServer(http.Dir("/public"))
// /*
// 	その後、マルチプレクサのHandle関数にハンドラを渡しています
// 	/static/から始まる全てのリクエストURLについて、StripPrefixでURLから/static/を取り除いた値のファイルがあれば、返信する
// 	<hルート>/img/a.img
// */
// mux.Handle("/static/", http.StripPrefix("/static/", files))

// /*
// 	第一引数にURL、第二引数はハンドラ関数名。ここでは、ルートがURLで指定された場合は、index関数を実行する
// 	パッケージmainがあるディレクトリにあるファイルはインクルードされる。
// 	import的な感じの書く必要ない
// */
// mux.HandleFunc("/", index)
// // error
// mux.HandleFunc("/err", err)

// // defined in route_auth.go
// mux.HandleFunc("/login", login)
// mux.HandleFunc("/logout", logout)
// mux.HandleFunc("/signup", signup)
// mux.HandleFunc("/signup_account", signupAccount)
// mux.HandleFunc("/authenticate", authenticate)

// // defined in route_thread.go
// mux.HandleFunc("/thread/new", newThread)
// mux.HandleFunc("/thread/create", createThread)
// mux.HandleFunc("/thread/post", postThread)
// mux.HandleFunc("/thread/read", readThread)

// server := &http.Server{
// 	Addr:    "0.0.0.0:8080",
// 	Handler: mux,
// }

// server.ListenAndServe()
// }
