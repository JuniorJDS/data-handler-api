package entity

import (
	"strconv"
	"strings"
)

const NULLVALUE = "NULL"

type UserData struct {
	CPF                  string
	Private              bool
	Incompleto           bool
	DataDaUltimaCompra   string
	TicketMedio          *float64
	TicketDaUltimaCompra *float64
	LojaMaisFrequente    string
	LojaDaUltimaCompra   string
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
	isPrivate, err := stringToBool(private)
	if err != nil {
		return nil, err
	}

	isIncompleto, err := stringToBool(incompleto)
	if err != nil {
		return nil, err
	}

	formattedDataDaUltimaCompra := stringToValidString(dataDaUltimaCompra)

	formattedTicketMedio, err := stringToFloat(ticketMedio)
	if err != nil {
		return nil, err
	}

	formattedTicketDaUltimaCompra, err := stringToFloat(ticketDaUltimaCompra)
	if err != nil {
		return nil, err
	}

	formattedLojaMaisFrequente := stringToValidString(lojaMaisFrequente)
	formattedLojaDaUltimaCompra := stringToValidString(lojaDaUltimaCompra)

	return &UserData{
		CPF:                  cpf,
		Private:              isPrivate,
		Incompleto:           isIncompleto,
		DataDaUltimaCompra:   formattedDataDaUltimaCompra,
		TicketMedio:          formattedTicketMedio,
		TicketDaUltimaCompra: formattedTicketDaUltimaCompra,
		LojaMaisFrequente:    formattedLojaMaisFrequente,
		LojaDaUltimaCompra:   formattedLojaDaUltimaCompra,
	}, nil
}

// TODO: levar para um utils
func stringToBool(data string) (bool, error) {
	var isData bool

	isData, err := strconv.ParseBool(data)
	if err != nil {
		return false, err
	}

	return isData, err
}

func stringToValidString(data string) string {
	if data == NULLVALUE {
		return ""
	}

	return data
}

func stringToFloat(data string) (*float64, error) {
	var floatData float64

	if data == NULLVALUE {
		return nil, nil
	}

	data = strings.Replace(data, ",", ".", 1)
	floatData, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return nil, err
	}

	return &floatData, nil
}
