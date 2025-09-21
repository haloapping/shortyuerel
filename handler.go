package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service
}

func NewHandler(s Service) Handler {
	return Handler{
		Service: s,
	}
}

func (h Handler) CreateShortUrl(c echo.Context) error {
	var req UrlReq
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"message": "invalid JSON",
			},
		)
	}

	u, err := h.Service.CreateShortUrl(c, req)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusBadRequest,
		map[string]any{
			"message": "short url is created",
			"data":    u,
		},
	)
}

func (h Handler) GetShortUrlByLongUrl(c echo.Context) error {
	var req UrlReq
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"message": "invalid JSON",
			},
		)
	}

	u, err := h.Service.CreateShortUrl(c, req)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(
		http.StatusBadRequest,
		map[string]any{
			"message": "short url is success retrieved",
			"data":    u,
		},
	)
}
