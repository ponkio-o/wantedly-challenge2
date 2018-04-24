package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

type Data struct {
	Message  string `json:"message"`
}

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", getIndex).Methods("GET")
    // 8080ポートで起動
    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}

func getIndex(w http.ResponseWriter, r *http.Request){
    // 構造体定義
    var data = Data{}
    data.Message = "Hello World!!"

    // jsonエンコード
    outputJson, err := json.Marshal(&data)
    if err != nil {
        panic(err)
    }

    // jsonヘッダーを出力
    w.Header().Set("Content-Type", "application/json")

    // jsonデータを出力
    fmt.Fprint(w, string(outputJson))
}
