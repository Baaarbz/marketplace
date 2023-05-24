package ad

import (
	. "barbz.dev/marketplace/internal/pkg/domain/ad"
	. "barbz.dev/marketplace/pkg/valueobject"
	"context"
	"database/sql"
	"github.com/google/uuid"
	"time"
)

const queryInsertAd = "INSERT INTO ads (id, title, description, price, postedat) VALUES ($1, $2, $3, $4, $5)"

type PostgresRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

func NewPostgresRepository(db *sql.DB, dbTimeout time.Duration) AdRepository {
	return &PostgresRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

func (repository *PostgresRepository) SaveAd(_ context.Context, ad Ad) (Ad, error) {
	var id, _ = uuid.NewUUID()
	var adId, _ = NewId(id.String())
	ad.SetId(adId)

	_, err := repository.db.Exec(queryInsertAd, ad.GetId().String(), ad.Title, ad.Description, ad.Price, ad.GetDate())
	return ad, err
}

func (repository *PostgresRepository) FindAdById(_ context.Context, id AdId) (Ad, error) {
	//for _, ad := range repository.ads {
	//	if ad.GetId() == id {
	//		return ad, nil
	//	}
	//}
	return Ad{}, nil
}

func (repository *PostgresRepository) FindAllAds(_ context.Context) (adResponse []Ad, err error) {
	//if len(repository.ads) < 5 {
	//	adResponse = repository.ads
	//} else {
	//	adResponse = repository.ads[:5]
	//}
	return
}
