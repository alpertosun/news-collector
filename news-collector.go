package main

import (
	"fmt"
	"github.com/alpertosun/news-collector/pkg/cmodels"
	"github.com/alpertosun/news-collector/pkg/getters"
	"github.com/alpertosun/news-collector/pkg/setters"
)

func main() {

	apiKey := "yourapikey"

	kategoriler := []string{
		"",
		"business",
		"entertainment",
		"health",
		"science",
		"sports",
		"technology",
	}

	for _,kategori := range kategoriler {

		fullData := new(cmodels.NewsAPIFormat)
		var data map[string]cmodels.Articles
		data = make(map[string]cmodels.Articles)

		err := getters.GetDataNewsAPI(kategori, apiKey, fullData)
		if err != nil {
			fmt.Println(err)
		}
		if kategori == "" {
			kategori = "gundem"
		}
		getters.GetDataFirebase("haber/kategori/" + kategori + "/articles",&data)


		setters.SetterMainFunc(fullData,data,kategori)
	}

}
