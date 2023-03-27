package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test title", Desc: "Test desc", Content: "Test content"},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	// json.NewEncoder関数を使用して、JSONエンコーダーを作成し、Encodeメソッドを使用して、オブジェクトをJSON形式に変換して、エンコーダーに書き込みます。
	json.NewEncoder(w).Encode(articles)
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST endpoint worked")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	// レスポンスの内容を http.ResponseWriter に書き込むときには、io.WriteString や fmt.Fprint、fmt.Fprintf などを使用できます
	// これは、http.ResponseWriter が io.Writer インタフェースを実装しているから
	fmt.Fprintf(w, "Homepage Endpoint Hit")
	// Fprintfは、第一引数に出力先のio.Writerを指定する必要があります。一方、printfは、標準出力に書き込みます。具体的には、os.Stdoutに書き込むことになります。
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true) // StrictSlash: /exampleというURIにアクセスした場合に、/example/にリダイレクトするようになります。

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	// log.Fatal関数で囲んでいるのは、エラーが発生した場合にその内容を出力してから終了するため
	// 第2引数のnil...http.DefaultServeMuxオブジェクトを指定するための引数
		// http.HandleFuncを使用してハンドラー関数を登録した場合には、nilを指定することができます
		// DefaultServeMux...Go言語の標準ライブラリのhttpパッケージにあるデフォルトのマルチプレクサであり、複数のハンドラーを登録して、リクエストのパスに応じて適切なハンドラーを呼び出すことができます。
		// マルチプレクサ...複数の入力を一つの出力にマッピングする装置やプログラムのことを指します。
		// log.Fatal(http.ListenAndServe(":8081", nil))
	log.Fatal(http.ListenAndServe(":8081", myRouter)) // httpをmyRouterに変更したので、第2引数がmyRouterになった
}

func main() {
	handleRequests()
}
