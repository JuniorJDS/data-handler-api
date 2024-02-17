package e2e_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

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

func TestUploadData__Simu__ExpectedSuccess(t *testing.T) {
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
