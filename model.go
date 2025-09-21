package main

import "github.com/guregu/null/v6"

type Url struct {
	Id        string    `json:"id"`
	LongUrl   string    `json:"longUrl"`
	ShortUrl  string    `json:"shortUrl"`
	CreatedAt null.Time `json:"createdAt"`
	UpdatedAt null.Time `json:"updatedAt"`
}

type UrlReq struct {
	LongUrl string `json:"longUrl"`
}
