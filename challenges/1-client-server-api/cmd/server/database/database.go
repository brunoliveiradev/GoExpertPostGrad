package database

import (
	"context"
	"database/sql"
	"github.com/brunoliveiradev/GoExpertPostGrad/challenges/1-client-server-api/pkg/domain"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	db   *sql.DB
	stmt *sql.Stmt
	mu   sync.Mutex
)

func InitSqliteDB() error {
	var err error

	dbPath := "server/database/database.db"
	dbDir := filepath.Dir(dbPath)

	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		err = os.MkdirAll(dbDir, 0755)
		if err != nil {
			return err
		}
	}

	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS currency_bids (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	code TEXT,
    	codein TEXT,
    	bid TEXT NOT NULL,
    	ask TEXT,
    	timestamp TEXT
  	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return err
	}

	stmt, err = db.Prepare(`INSERT INTO currency_bids (code, codein, bid, ask, timestamp) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}

	return nil
}

func CloseSqliteDB() error {
	mu.Lock()
	defer mu.Unlock()

	var stmtErr, dbErr error
	if stmt != nil {
		stmtErr = stmt.Close()
	}
	if db != nil {
		dbErr = db.Close()
	}

	if stmtErr != nil {
		return stmtErr
	}
	if dbErr != nil {
		return dbErr
	}

	return nil
}

func SaveCurrencyInfo(ctx context.Context, currencyInfo *domain.CurrencyInfo) error {
	mu.Lock()
	defer mu.Unlock()

	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.Stmt(stmt).ExecContext(ctx, currencyInfo.Code, currencyInfo.Codein, currencyInfo.Bid, currencyInfo.Ask, currencyInfo.Timestamp)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
