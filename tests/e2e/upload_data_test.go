package e2e_test

import (
	"context"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

func listData(b *BaseTest) ([]*UserData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := "SELECT cpf, private, incompleto, datadaultimacompra, ticketmedio, ticketdaultimacompra, lojamaisfrequente, lojadaultimacompra, isvalidcpforcnpj FROM userdata"

	rows, err := b.database.Query(ctx, query)
	if err != nil {
		log.Printf("Failed query models: %v\n", err)
		return nil, err
	}

	fetchUserData := []*UserData{}
	for rows.Next() {
		var m UserData
		err := rows.Scan(
			&m.CPF, &m.Private, &m.Incompleto, &m.DataDaUltimaCompra, &m.TicketMedio,
			&m.TicketDaUltimaCompra, &m.LojaMaisFrequente, &m.LojaDaUltimaCompra, &m.IsValidCPForCNPJ,
		)
		if err != nil {
			log.Printf("Error to fetch model: %v\n", err)
			return nil, err
		}

		fetchUserData = append(fetchUserData, &m)
	}
	if err != nil {
		return nil, err
	}
	return fetchUserData, nil
}

func countRows(b *BaseTest) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := "SELECT COUNT(1) FROM userdata"

	var count int
	err := b.database.QueryRow(ctx, query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func TestUploadData__UploadProcessAndPersistTheBaseTestData__ExpectedSuccess(t *testing.T) {
	baseTest := NewBaseTest()
	defer baseTest.TearDownTest()

	// FIXTURE
	filePath := "./data/base_teste.txt"
	payload, writer, _ := baseTest.createFileMultipartPayload(filePath)

	// EXERCISE
	url := "/api/v1/upload"
	resp, errResp := baseTest.appClient("POST", url, payload, writer)

	// ASSERTS
	assert.NoError(t, errResp)

	assert.Equal(t, http.StatusCreated, resp.Code)

	numberOfRows, _ := countRows(baseTest)
	assert.Equal(t, 49998, numberOfRows)
}

func TestUploadData__UploadProcessAndPersistASmallBaseTestWithInvalidCPF__ExpectedSuccess(t *testing.T) {
	baseTest := NewBaseTest()
	defer baseTest.TearDownTest()

	// FIXTURE
	filePath := "./data/small_base_test_with_false_cpf.txt"
	payload, writer, _ := baseTest.createFileMultipartPayload(filePath)

	// EXERCISE
	url := "/api/v1/upload"
	resp, errResp := baseTest.appClient("POST", url, payload, writer)

	// ASSERTS
	assert.NoError(t, errResp)

	assert.Equal(t, http.StatusCreated, resp.Code)

	numberOfRows, _ := countRows(baseTest)
	assert.Equal(t, 4, numberOfRows)

	resultData, _ := listData(baseTest)

	ticketValue := float64(335.38)
	parsedDate, _ := time.Parse("2006-01-02", "2010-01-13")
	expectedData := []*UserData{
		{
			CPF:                  "00000000000",
			Private:              false,
			Incompleto:           false,
			DataDaUltimaCompra:   parsedDate,
			TicketMedio:          &ticketValue,
			TicketDaUltimaCompra: &ticketValue,
			LojaMaisFrequente:    "79379491000183",
			LojaDaUltimaCompra:   "79379491000183",
			IsValidCPForCNPJ:     false,
		},
		{
			CPF:                  "78932742977",
			Private:              true,
			Incompleto:           false,
			DataDaUltimaCompra:   parsedDate,
			TicketMedio:          &ticketValue,
			TicketDaUltimaCompra: &ticketValue,
			LojaMaisFrequente:    "79379491000183",
			LojaDaUltimaCompra:   "79379491000183",
			IsValidCPForCNPJ:     false,
		},
		{
			CPF:                  "06029876822",
			Private:              false,
			Incompleto:           true,
			DataDaUltimaCompra:   parsedDate,
			TicketMedio:          &ticketValue,
			TicketDaUltimaCompra: &ticketValue,
			LojaMaisFrequente:    "79379491000183",
			LojaDaUltimaCompra:   "79379491000183",
			IsValidCPForCNPJ:     false,
		},
		{
			CPF:                  "04109164125",
			Private:              false,
			Incompleto:           false,
			DataDaUltimaCompra:   parsedDate,
			TicketMedio:          &ticketValue,
			TicketDaUltimaCompra: &ticketValue,
			LojaMaisFrequente:    "79379491000183",
			LojaDaUltimaCompra:   "79379491000183",
			IsValidCPForCNPJ:     true,
		},
	}

	assert.ElementsMatch(t, expectedData, resultData)
}

func TestUploadData__TryToUploadDataWithoutFileMultipart__ExpectedBadRequest(t *testing.T) {
	baseTest := NewBaseTest()
	defer baseTest.TearDownTest()

	// EXERCISE
	url := "/api/v1/upload"
	resp, errResp := baseTest.appClient("POST", url, nil, nil)

	// ASSERTS
	assert.NoError(t, errResp)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	numberOfRows, _ := countRows(baseTest)
	assert.Equal(t, 0, numberOfRows)
}
