package service

import (
	"bufio"
	"io"
	"strings"
	"sync"

	"github.com/JuniorJDS/data-handler-api/entity"
	"github.com/JuniorJDS/data-handler-api/repository"
)

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
	tokens := make(chan struct{}, 50)

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

	err := fh.UserRepository.InsertManyRows(rows)
	if err != nil {
		return err
	}
	return nil
}
