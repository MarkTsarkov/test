package repo

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/marktsarkov/test/model"
)

type repo struct {
	db *pgxpool.Pool
}

func NewClickRepo(db *pgxpool.Pool) Irepo {
	return &repo{db: db}
}

func (r *repo) SaveClicks(ctx context.Context, data map[model.Click]int) error {
	var (
		placeholders []string
		args         []interface{}
	)
	elemId := 1
	for click, count := range data {
		placeholders = append(placeholders,
			fmt.Sprintf("($%d, $%d, $%d)", elemId, elemId+1, elemId+2))
		args = append(args, click.BannerID, click.Ts, count)
		elemId += 3
	}

	queryAdd := fmt.Sprintf("INSERT INTO click_stats (banner_id, ts, count) VALUES "+
		"%s ON CONFLICT (banner_id, ts) "+
		"DO UPDATE SET count = click_stats.count + EXCLUDED.count;",
		strings.Join(placeholders, ", "))

	result, err := r.db.Exec(ctx, queryAdd, args...)
	if err != nil {
		log.Printf("error while db exec: %v", err)
		return err
	}
	rowsAffected := result.RowsAffected()
	log.Printf("rows affected: %d\n", rowsAffected)
	return nil
}

func (r *repo) GetStats(ctx context.Context, bannerID int, tsFrom, tsTo time.Time) ([]model.ClickStat, error) {
	return nil, nil
}
