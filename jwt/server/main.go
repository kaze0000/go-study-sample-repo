package main

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("mysupersecretphrase")

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "super secret information")
}

// super secret informationを特定の人しか見られないようにする
// endpointという関数を引数にもつ
// http.Handlerは、HTTPリクエストを処理するためのインターフェースであり、ServeHTTPというメソッドを持つ。
func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	// funv isAuthorizedはmiddleware functionである
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error){
				// interface{}型の値：JWTトークンの署名キーを表す任意の型の値を返す(nilやmySigningKeyに対応できる)
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					// 検証を行うサーバー側であらかじめアルゴリズムを指定しておく alg:none攻撃対策
					// 詳細 https://scgajge12.hatenablog.com/entry/jwt_security
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				 endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}

func handleRequests() {
	// http.HandleFunc("/", homePage) //* http.HandleFuncとhttp.Handleでは、ハンドラーの実装方法が異なる
	http.Handle("/", isAuthorized(homePage))

	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	fmt.Println("my simple server")
	handleRequests()
}
