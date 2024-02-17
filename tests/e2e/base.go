package e2e_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/JuniorJDS/data-handler-api/api"
	"github.com/JuniorJDS/data-handler-api/infra"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"
)

type BaseTest struct {
	suite.Suite
	database *pgxpool.Pool
}

func NewBaseTest() *BaseTest {
	return &BaseTest{
		database: infra.GetDB(),
	}
}

func (b *BaseTest) appClient(verbHTTP, pathEndpoint string, body io.Reader, writer *multipart.Writer) (*httptest.ResponseRecorder, error) {
	request, err := http.NewRequest(verbHTTP, pathEndpoint, body)
	if err != nil {
		fmt.Println("Erro to make request: ", err.Error())
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	if writer != nil {
		request.Header.Set("Content-Type", writer.FormDataContentType())
	}

	response := httptest.NewRecorder()

	handler := api.HttpHandler()
	handler.ServeHTTP(response, request)

	return response, nil
}

func (b *BaseTest) TearDownTest() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := fmt.Sprintf("TRUNCATE TABLE %s;", "userdata")

	_, err := b.database.Exec(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func (b *BaseTest) createFileMultipartPayload(filePath string) (*bytes.Buffer, *multipart.Writer, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		return nil, nil, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, nil, err
	}
	writer.Close()

	return payload, writer, nil
}
