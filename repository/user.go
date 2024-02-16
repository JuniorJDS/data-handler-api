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

func (u *UserRepository) InsertManyRows(data <-chan *entity.UserData) error {
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
	for d := range data {
		cpfAux = append(cpfAux, d.CPF)
		privateAux = append(privateAux, d.Private)
		incompletoAux = append(incompletoAux, d.Incompleto)
		datadaultimacompraAux = append(datadaultimacompraAux, d.DataDaUltimaCompra)
		ticketmedioAux = append(ticketmedioAux, d.TicketMedio)
		ticketdaultimacompraAux = append(ticketdaultimacompraAux, d.TicketDaUltimaCompra)
		lojamaisfrequenteAux = append(lojamaisfrequenteAux, d.LojaMaisFrequente)
		lojadaultimacompraAux = append(lojadaultimacompraAux, d.LojaDaUltimaCompra)
		isValidCPForCNPJ = append(isValidCPForCNPJ, d.IsValidCPForCNPJ)
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
