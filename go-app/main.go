package main

import (
	"io"
	"encoding/json"
	"fmt"
	"net/http"
	"book-recom/models"
)

func main() {

	db := models.ConnectDB()

	models.GetData(db)

	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8000", nil)
}


func handleRequest(w http.ResponseWriter, r *http.Request) {
	// 共通処理
	fmt.Println("Common process")

	// URLに応じたハンドラを実行
	switch r.URL.Path {
	case "/xxx":
		handleXxx(w, r)
	default:
		http.NotFound(w, r)
	}
}



func handleXxx(w http.ResponseWriter, r *http.Request) {
	// inputDataの構造体を用意
	type requestParam struct {
		Text string `json:"text"`
		Id int `json:"id"`
	}

	inputData := &requestParam{}

	// ポインタ型で渡すこと
	err := ConstructInputData(r, inputData)

	fmt.Println(inputData.Text, inputData.Id)

	data := map[string]string{
		"message": "Hello, JSON!",
	}

	w.Header().Set("Content-Type", "application/json")
	
	// jsonに変換する
	res, err := json.Marshal(data)


	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	}

	// レスポンスを返却する
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}



func ConstructInputData[T any](r *http.Request, inputData T) error {
	// リクエストを読み込む。
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	// リクエストを引数に受け取った構造体にマッピングする
	err = json.Unmarshal(body, inputData)
	if err != nil {
		return err
	}

	return nil
}