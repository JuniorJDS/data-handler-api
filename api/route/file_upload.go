package route

import (
	"net/http"

	"github.com/JuniorJDS/data-handler-api/service"
)

type FileUploadRoute struct {
	FileService service.FileHandler
}

func NewFileUploadRoute() *FileUploadRoute {
	return &FileUploadRoute{}
}

func (fr *FileUploadRoute) UploadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("file")
	if err != nil {
		return
	}

	defer file.Close()

	// path := fmt.Sprintf("/%s", fileHeader.Filename)
	err = fr.FileService.Process(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	message := map[string]string{
		"msg": "Dados atualizados com Sucesso.",
	}
	responseWithJSON(w, http.StatusCreated, message)
}