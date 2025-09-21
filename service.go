package main

import (
	"github.com/labstack/echo/v4"
)

type Service struct {
	Repository
}

func NewService(r Repository) Service {
	return Service{
		Repository: r,
	}
}

func (s Service) CreateShortUrl(c echo.Context, req UrlReq) (Url, error) {
	u, err := s.Repository.CreateShortUrl(c, req)
	if err != nil {
		return Url{}, err
	}

	return u, nil
}

func (s Service) GetShortUrlByLongUrl(c echo.Context, req UrlReq) (Url, error) {
	u, err := s.Repository.GetShortUrlByLongUrl(c, req)
	if err != nil {
		return Url{}, err
	}

	return u, nil
}
