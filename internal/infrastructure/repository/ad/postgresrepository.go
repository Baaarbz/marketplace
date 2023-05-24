package ad

import (
	. "barbz.dev/marketplace/internal/pkg/domain/ad"
	. "barbz.dev/marketplace/pkg/valueobject"
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"time"
)

const queryInsertAd = "INSERT INTO ads (id, title, description, price, postedat) VALUES ($1, $2, $3, $4, $5)"
const queryFindAdById = "SELECT id, title, description, price, postedat FROM ads WHERE id = $1"

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
	ad.Id = adId

	_, err := repository.db.Exec(queryInsertAd, ad.Id, ad.Title, ad.Description, fmt.Sprintf("%.2f", ad.Price), ad.Date)
	return ad, err
}

func (repository *PostgresRepository) FindAdById(_ context.Context, id AdId) (Ad, error) {
	row := repository.db.QueryRow(queryFindAdById, id)

	ad := Ad{}
	err := row.Scan(&ad.Id, &ad.Title, &ad.Description, &ad.Price, &ad.Date)

	return ad, err
}

func (repository *PostgresRepository) FindAllAds(_ context.Context) (adResponse []Ad, err error) {
	//if len(repository.ads) < 5 {
	//	adResponse = repository.ads
	//} else {
	//	adResponse = repository.ads[:5]
	//}
	return
}
