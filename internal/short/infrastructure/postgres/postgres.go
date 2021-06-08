package postgres

import (
	"github.com/bearname/url-short/internal/common/uuid"
	"github.com/bearname/url-short/internal/short/app"
	"github.com/bearname/url-short/internal/short/domain"
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"time"
)

const errUniqueConstraint = "23505"

type rawUrl struct {
	Id           string    `db:"id"`
	OriginalUrl  string    `db:"original_url"`
	CreationDate time.Time `db:"creation_date"`
	Alias        string    `db:"alias"`
}

func (r *UrlRepositoryImpl) NextID() domain.UrlID {
	return domain.UrlID(uuid.Generate())
}

type UrlRepositoryImpl struct {
	connPool *pgx.ConnPool
}

func NewUrlRepository(connPool *pgx.ConnPool) *UrlRepositoryImpl {
	u := new(UrlRepositoryImpl)
	u.connPool = connPool
	return u
}

func (r *UrlRepositoryImpl) Create(item domain.Url) error {
	_, err := r.connPool.Exec(
		`INSERT INTO urls (id, original_url, alias) 
			 VALUES ($1, $2, $3)`,
		item.Id.String(),
		item.OriginalUrl,
		item.Alias)

	if err != nil {
		pgErr, ok := err.(pgx.PgError)
		log.Info(pgErr.Code)
		if ok && pgErr.Code == errUniqueConstraint {
			return app.ErrDuplicateUrl
		}
		return errors.WithStack(err)
	}

	return nil
}

func (r *UrlRepositoryImpl) FindByAlias(alias string) (*domain.Url, error) {
	var raw rawUrl
	query := `SELECT id, original_url, alias
              FROM urls WHERE alias = $1`
	err := r.connPool.QueryRow(query, alias).Scan(
		&raw.Id,
		&raw.OriginalUrl,
		&raw.Alias)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = app.ErrUrlNotFound
		}
		return nil, errors.WithStack(err)
	}

	return mapToUrl(raw)
}

func mapToUrl(raw rawUrl) (*domain.Url, error) {
	itemID, _ := uuid.FromString(raw.Id)
	item := &domain.Url{
		Id:          domain.UrlID(itemID),
		OriginalUrl: raw.OriginalUrl,
		Alias:       raw.Alias,
	}

	return item, nil
}
