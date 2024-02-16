package repository

import (
	"context"
	"time"

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
	query := `
		INSERT INTO userdata
		(cpf, private, incompleto, datadaultimacompra, ticketmedio, ticketdaultimacompra, lojamaisfrequente, lojadaultimacompra, isvalidcpforcnpj) 
		  (
			select * from unnest($1::varchar[], $2::boolean[], $3::boolean[], $4::date[], $5::numeric[], $6::numeric[], $7::varchar[], $8::varchar[], $9::boolean[])
		  )`

	cpfAux, privateAux, incompletoAux := []string{}, []bool{}, []bool{}
	datadaultimacompraAux, ticketmedioAux, ticketdaultimacompraAux := []time.Time{}, []*float64{}, []*float64{}
	lojamaisfrequenteAux, lojadaultimacompraAux := []string{}, []string{}
	isValidCPForCNPJ := []bool{}
	for _, userData := range data {
		cpfAux = append(cpfAux, userData.CPF)
		privateAux = append(privateAux, userData.Private)
		incompletoAux = append(incompletoAux, userData.Incompleto)
		datadaultimacompraAux = append(datadaultimacompraAux, userData.DataDaUltimaCompra)
		ticketmedioAux = append(ticketmedioAux, userData.TicketMedio)
		ticketdaultimacompraAux = append(ticketdaultimacompraAux, userData.TicketDaUltimaCompra)
		lojamaisfrequenteAux = append(lojamaisfrequenteAux, userData.LojaMaisFrequente)
		lojadaultimacompraAux = append(lojadaultimacompraAux, userData.LojaDaUltimaCompra)
		isValidCPForCNPJ = append(isValidCPForCNPJ, userData.IsValidCPForCNPJ)
	}

	if _, err := u.db.Exec(
		context.Background(),
		query,
		cpfAux,
		privateAux,
		incompletoAux,
		datadaultimacompraAux,
		ticketmedioAux,
		ticketdaultimacompraAux,
		lojamaisfrequenteAux,
		lojadaultimacompraAux,
		isValidCPForCNPJ,
	); err != nil {
		return err
	}

	return nil
}
