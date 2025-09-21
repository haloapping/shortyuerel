package main

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/oklog/ulid/v2"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewReposity(p *pgxpool.Pool) Repository {
	return Repository{
		pool: p,
	}
}

func (r Repository) CreateShortUrl(c echo.Context, req UrlReq) (Url, error) {
	ctx := c.Request().Context()
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return Url{}, nil
	}

	q := `
		INSERT INTO urls(id, long_url, short_url)
		VALUES($1, $2, $3)
		RETURNING *;
	`
	shortId, err := gonanoid.Generate("abcdefghijklmnopqrstuvwxyz0123456789_", 10)
	shortUrl := fmt.Sprintf("https://yuerel%s", shortId)
	row := tx.QueryRow(ctx, q, ulid.Make().String(), req.LongUrl, shortUrl)
	var u Url
	err = row.Scan(&u.Id, &u.LongUrl, &u.ShortUrl, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		tx.Rollback(ctx)

		return Url{}, err
	}
	tx.Commit(ctx)

	return u, nil
}

func (r Repository) GetShortUrlByLongUrl(c echo.Context, req UrlReq) (Url, error) {
	ctx := c.Request().Context()
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return Url{}, nil
	}

	q := `
		SELECT *
		FROM urls
		WHERE long_url = $1;
	`
	row := tx.QueryRow(ctx, q, req.LongUrl)
	var u Url
	err = row.Scan(&u.Id, &u.LongUrl, &u.ShortUrl, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		tx.Rollback(ctx)

		return Url{}, err
	}
	tx.Commit(ctx)

	return u, nil
}
