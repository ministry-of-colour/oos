package api

import (
	"database/sql"
	"os"
	"time"

	dat "github.com/helloeave/dat/dat"
	runner "github.com/helloeave/dat/sqlx-runner"
	"github.com/sirupsen/logrus"
)

type DB interface {
	StockReport(brand string) (*stockReport, error)
}

type stock struct {
	Brand string `db:"brand"`
	Item  string `db:"id"`
	Descr string `db:"name"`
	Qty   string `db:"qty"`
}

type stockReport struct {
	Note  string
	Items []*stock
}

type pgDB struct {
	log *logrus.Logger
	db  *runner.DB
}

// StockReport retuns a struture with all the stock data needed
func (h *pgDB) StockReport(brand string) (*stockReport, error) {
	retval := &stockReport{}

	// Get the overall status
	err := h.db.Select("descr").From("brandStockStatus").Limit(1).QueryScalar(&retval.Note)
	if err != nil {
		return retval, err
	}

	// Get the items part
	err = h.db.Select("brand,id,name,round(random()*100) as qty").From("items").Where("brand=$1", brand).QueryStructs(&retval.Items)
	return retval, err
}

// newDB is a private function to setup the database connection
func newDB(l *logrus.Logger) *pgDB {

	// create a normal database connection through database/sql
	db, err := sql.Open("postgres", os.Getenv("API_DSN"))
	if err != nil {
		l.WithError(err).Fatal("could not connect to DB")
	}

	// ensures the database can be pinged with an exponential backoff (15 min)
	runner.MustPing(db)

	// set to reasonable values for production
	db.SetMaxIdleConns(4)
	db.SetMaxOpenConns(16)

	// set this to enable interpolation
	dat.EnableInterpolation = true

	// set to check things like sessions closing.
	// Should be disabled in production/release builds.
	dat.Strict = false

	// Log any query over 10ms as warnings. (optional)
	runner.LogQueriesThreshold = 10 * time.Millisecond
	l.Info("Connected to DB")
	return &pgDB{
		log: l,
		db:  runner.NewDB(db, "postgres"),
	}
}
