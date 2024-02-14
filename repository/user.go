package repository

import (
	"fmt"

	"github.com/JuniorJDS/data-handler-api/entity"
	"github.com/JuniorJDS/data-handler-api/infra"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: infra.GetDB(),
	}
}

func (u *UserRepository) InsertManyRows(data []entity.UserData) error {
	sqlQuery := `INSERT INTO 
		userData (cpf, private, incompleto, lastdate, avgticket, lastTicket,  storefrequent, storelast) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	// u.db.Exec()
	fmt.Println(sqlQuery)

	return nil
}
