package entity

import (
	"time"

	"github.com/JuniorJDS/data-handler-api/utils"
)

type UserData struct {
	CPF                  string
	Private              bool
	Incompleto           bool
	DataDaUltimaCompra   time.Time
	TicketMedio          *float64
	TicketDaUltimaCompra *float64
	LojaMaisFrequente    string
	LojaDaUltimaCompra   string
	IsValidCPForCNPJ     bool
}

func NewUserData(
	cpf,
	private,
	incompleto,
	dataDaUltimaCompra,
	ticketMedio,
	ticketDaUltimaCompra,
	lojaMaisFrequente,
	lojaDaUltimaCompra string,
) (*UserData, error) {
	formattedCPFCNPJ, isValidCPForCNPJ := utils.FormatCPFAndCNPJ(cpf)

	isPrivate, err := utils.StringToBool(private)
	if err != nil {
		return nil, err
	}

	isIncompleto, err := utils.StringToBool(incompleto)
	if err != nil {
		return nil, err
	}

	formattedDataDaUltimaCompra, err := utils.StringToDate(dataDaUltimaCompra)
	if err != nil {
		return nil, err
	}

	formattedTicketMedio, err := utils.StringToFloat(ticketMedio)
	if err != nil {
		return nil, err
	}

	formattedTicketDaUltimaCompra, err := utils.StringToFloat(ticketDaUltimaCompra)
	if err != nil {
		return nil, err
	}

	formattedLojaMaisFrequente, _ := utils.FormatCPFAndCNPJ(lojaMaisFrequente)
	formattedLojaDaUltimaCompra, _ := utils.FormatCPFAndCNPJ(lojaDaUltimaCompra)

	return &UserData{
		CPF:                  formattedCPFCNPJ,
		Private:              isPrivate,
		Incompleto:           isIncompleto,
		DataDaUltimaCompra:   formattedDataDaUltimaCompra,
		TicketMedio:          formattedTicketMedio,
		TicketDaUltimaCompra: formattedTicketDaUltimaCompra,
		LojaMaisFrequente:    formattedLojaMaisFrequente,
		LojaDaUltimaCompra:   formattedLojaDaUltimaCompra,
		IsValidCPForCNPJ:     isValidCPForCNPJ,
	}, nil
}
