package getters

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/alpertosun/news-collector/pkg/firebaseconn"
	"time"
)

func init()  {

}

func GetDataFirebase(path string,objekt interface{}) interface{} {
	firebaseDB, ctx := firebaseconn.CreateNewConnection()
	err := firebaseDB.NewRef(path).Get(ctx,&objekt)
	if err != nil {
		fmt.Println("asd",err)
	}
	return objekt
}

func GetDataNewsAPI(kategori string, apiKey string, objekt interface{}) error {
	NewApiUrl := "http://newsapi.org/v2/top-headlines?country=tr&apiKey=" + apiKey
	if kategori != "" {
		NewApiUrl = "http://newsapi.org/v2/top-headlines?country=tr&category="+ kategori + "&apiKey=" + apiKey
	}
	var myClient = &http.Client{Timeout: 10 * time.Second}
	r, err := myClient.Get(NewApiUrl)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(objekt)
}