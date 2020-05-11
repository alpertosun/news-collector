package cmodels

import "time"

type
	Articles struct {
		Source struct {
			ID   interface{} `json:"id"`
			Name string      `json:"name"`
		} `json:"source"`
		Category    string      `json:"category"`
		Author      string      `json:"author"`
		UniqueName  string      `json:"uname"`
		Title       string      `json:"title"`
		Description string      `json:"description"`
		URL         string      `json:"url"`
		URLToImage  string      `json:"urlToImage"`
		PublishedAt time.Time   `json:"publishedAt"`
		Content     interface{} `json:"content"`
	}

type NewsAPIFormat struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Category     string `json:"category"`
	Articles     []struct {
		Source struct {
			ID   interface{} `json:"id"`
			Name string      `json:"name"`
		} `json:"source"`
		Author      string      `json:"author"`
		UniqueName  string      `json:"uname"`
		Title       string      `json:"title"`
		Description string      `json:"description"`
		URL         string      `json:"url"`
		URLToImage  string      `json:"urlToImage"`
		PublishedAt time.Time   `json:"publishedAt"`
		Content     interface{} `json:"content"`
	} `json:"articles"`
}

type SiraNo struct {
	Gundem    string `json:"sira"`
	Business  string `json:"business"`
	Entertain string `json:"entertain"`
	Health    string `json:"health"`
	Science   string `json:"science"`
	Sports    string `json:"sports"`
	Tech      string `json:"tech"`
}
