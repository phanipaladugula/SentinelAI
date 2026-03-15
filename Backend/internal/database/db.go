package database

import(
	"fmt"

	"sentinel-ai/internal/camera/config"

	"github.com/jmoiron/sqlx"
	_"github.com/lib/pq"
)

func NewDB(cfg *config.Config)(*sqlx.DB,error){
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	db,err := sqlx.Connect("postgres",dsn)

	if err!=nil{
		return nil,err
	}

	return db,nil
}