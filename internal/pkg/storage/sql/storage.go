package sql

import (
	"fmt"
	"os"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	trmgorm "github.com/avito-tech/go-transaction-manager/gorm"
	"github.com/hse-revizor/analysis-service/internal/pkg/models"
	"github.com/hse-revizor/analysis-service/internal/utils/config"

	"github.com/stretchr/testify/require"
)

type Storage struct {
	db     *gorm.DB
	getter *trmgorm.CtxGetter
}

func New(db *gorm.DB, getter *trmgorm.CtxGetter) *Storage {
	return &Storage{
		db:     db,
		getter: getter,
	}
}

func buildDSN(cfg *config.Config) string {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s  dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.SSLMode)

	return dsn
}

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(buildDSN(cfg)), &gorm.Config{
		TranslateError: true,
	})
}

func MustNewPostgresDB(cfg *config.Config) *gorm.DB {
	db, err := NewPostgresDB(cfg)
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.Analyze{})
	if err != nil {
		panic(err)
	}

	return db
}

func NewSQLIteDB(cfg *config.Config) (*gorm.DB, error) {

	return gorm.Open(sqlite.Open("devdb.db"), &gorm.Config{
		TranslateError: true,
	})
}

func MustNewSQLite(cfg *config.Config) *gorm.DB {
	db, err := NewSQLIteDB(cfg)
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.Analyze{})

	if err != nil {
		panic(err)
	}
	return db
}

func MustNewTestDB(t *testing.T) *gorm.DB {
	const dbName = "test_storage.db"
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		require.NoError(t, err)
	}

	if err != nil {
		require.NoError(t, err)
	}

	err = db.AutoMigrate(&models.Analyze{})
	require.NoError(t, err)

	t.Cleanup(func() {
		dbInstance, err := db.DB()
		require.NoError(t, err)
		require.NoError(t, dbInstance.Close())
		require.NoError(t, os.Remove(dbName))
	})

	return db
}
