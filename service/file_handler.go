package service

import (
	"bufio"
	"io"
	"log"
	"strings"
	"sync"

	"github.com/JuniorJDS/data-handler-api/entity"
	"github.com/JuniorJDS/data-handler-api/repository"
)

const maxTokens = 50

type FileHandler struct {
	UserRepository repository.UserRepository
}

func NewFileHandler() *FileHandler {
	return &FileHandler{
		UserRepository: *repository.NewUserRepository(),
	}
}

func (fh *FileHandler) Process(file io.Reader) error {
	s := bufio.NewScanner(file)
	s.Scan() // skip first row

	var wg sync.WaitGroup

	rows := make(chan *entity.UserData)
	tokens := make(chan struct{}, maxTokens)

	for s.Scan() {
		wg.Add(1)
		tokens <- struct{}{}
		row := strings.Fields(s.Text())

		go func(row []string) {
			defer wg.Done()
			user, err := entity.NewUserData(
				row[0],
				row[1],
				row[2],
				row[3],
				row[4],
				row[5],
				row[6],
				row[7],
			)
			if err != nil {
				log.Printf("Error creating UserData: %v", err)
				return
			}

			<-tokens
			rows <- user
		}(row)
	}

	go func() {
		wg.Wait()
		close(rows)
	}()

	var result []entity.UserData
	for row := range rows {
		result = append(result, *row)
	}

	if err := fh.UserRepository.InsertManyRows(result); err != nil {
		return err
	}
	return nil
}
