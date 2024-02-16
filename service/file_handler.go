package service

import (
	"bufio"
	"fmt"
	"io"
	"strings"

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

	var userData []entity.UserData
	for s.Scan() {
		row := strings.Fields(s.Text())

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
			fmt.Println(err)
		}

		userData = append(userData, *user)
	}

	err := fh.UserRepository.InsertManyRows(userData)
	if err != nil {
		return err
	}
	return nil
}
