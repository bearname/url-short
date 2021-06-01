package infrastructure

import (
	"github.com/bearname/url-short/pkg/common/uuid"
	"github.com/bearname/url-short/pkg/short/app"
	"github.com/bearname/url-short/pkg/short/domain"
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
	"time"
)

const errUniqueConstraint = "23505"

type rawUrl struct {
	Id             string    `db:"id"`
	OriginalUrl    string    `db:"original_url"`
	CreationDate   time.Time `db:"creation_date"`
	ExpirationDate time.Time `db:"expiration_date"`
	CustomAlias    string    `db:"custom_alias"`
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
		`INSERT INTO urls (id, original_url, expiration_date, custom_alis) 
			 VALUES ($1, $2, $3, $4)`,
		item.Id.String(),
		item.OriginalUrl,
		item.ExpirationDate,
		item.CustomUrl)

	if err != nil {
		pgErr, ok := err.(pgx.PgError)
		if ok && pgErr.Code == errUniqueConstraint {
			return app.ErrDuplicateUrl
		}
		return errors.WithStack(err)
	}
	return nil
}

func (r *UrlRepositoryImpl) FindById(id domain.UrlID) (*domain.Url, error) {
	var raw rawUrl
	query := `SELECT id, original_url, expiration_date, custom_alias
              FROM urls WHERE id = $1`
	err := r.connPool.QueryRow(query, id.String()).Scan(
		&raw.Id,
		&raw.OriginalUrl,
		&raw.ExpirationDate,
		&raw.CustomAlias)
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
		Id:             domain.UrlID(itemID),
		CreationDate:   raw.CreationDate,
		ExpirationDate: raw.ExpirationDate,
		CustomUrl:      raw.CustomAlias,
	}

	return item, nil
}
