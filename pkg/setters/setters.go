package setters

import (
	"fmt"
	"github.com/alpertosun/news-collector/pkg/firebaseconn"
	"github.com/alpertosun/news-collector/pkg/getters"
	"github.com/alpertosun/news-collector/pkg/cmodels"
	"strconv"
	"strings"
)

func setData(path string,objekt interface{})  {
	firebaseDB, ctx := firebaseconn.CreateNewConnection()

	err := firebaseDB.NewRef(path).Set(ctx,objekt)
	if err != nil {
		fmt.Println(err)
	}

}

// yenisi objekt1 eskisi objekt2
func SetterMainFunc(objekt1 *cmodels.NewsAPIFormat, objekt2 map[string]cmodels.Articles, kategori string)  {

	olmayanlar, siraNo := isInInterface(objekt1,objekt2,kategori)



	for _,v := range olmayanlar.Articles {

		//sirano objesindeki kategori seçimi yüzünden switch kullandım
		switch kategori {
		case "gundem":
			siraNo.Gundem = incrementNumber(siraNo.Gundem)
			setData("haber/kategori/" + kategori + "/articles/" + siraNo.Gundem,v)
		case "business":
			siraNo.Business = incrementNumber(siraNo.Business)
			setData("haber/kategori/" + kategori + "/articles/" + siraNo.Business,v)
		case "technology":
			siraNo.Tech = incrementNumber(siraNo.Tech)
			setData("haber/kategori/" + kategori + "/articles/" + siraNo.Tech,v)
		case "sports":
			siraNo.Sports = incrementNumber(siraNo.Sports)
			setData("haber/kategori/" + kategori + "/articles/" + siraNo.Sports,v)
		case "science":
			siraNo.Science = incrementNumber(siraNo.Science)
			setData("haber/kategori/" + kategori + "/articles/" + siraNo.Science,v)
		case "health":
			siraNo.Health = incrementNumber(siraNo.Health)
			setData("haber/kategori/" + kategori + "/articles/" + siraNo.Health,v)
		case "entertainment":
			siraNo.Entertain = incrementNumber(siraNo.Entertain)
			setData("haber/kategori/" + kategori + "/articles/" + siraNo.Entertain,v)
		}

	}

	setData("haber/sira",siraNo)
	fmt.Println(len(olmayanlar.Articles),"haber",kategori,"kategorisine eklendi.")
}

func isInInterface(objekt1 *cmodels.NewsAPIFormat, objekt2 map[string]cmodels.Articles, kategori string) (*cmodels.NewsAPIFormat,cmodels.SiraNo) {

	olmayanlar := new(cmodels.NewsAPIFormat)

	var siraNo cmodels.SiraNo
	getters.GetDataFirebase("haber/sira",&siraNo)


	for _, v := range objekt1.Articles {

		var varMi bool = false

		for _, b := range objekt2 {
			if v.Title == b.Title {
				varMi = true
				break
			}
		}

		if varMi == false {

			olmayanlar.Articles = append(olmayanlar.Articles,v)
			olmayanlar.Category = kategori

		}
	}
	return olmayanlar,siraNo
}

func incrementNumber(string2 string) string {
	intnumber, _ := strconv.Atoi(strings.Replace(string2,"\"","",-1))
	intnumber++
	return strconv.Itoa(intnumber)
}