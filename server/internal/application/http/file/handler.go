package file

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	error_handler "github.com/PakSerVal/bear-masha/internal/application/http"
	"github.com/PakSerVal/bear-masha/internal/domain/file/usecase/upload"
)

type Handler struct {
	usecase upload.Usecase
}

func New(usecase upload.Usecase) *Handler {
	return &Handler{
		usecase: usecase,
	}
}

func (h *Handler) Upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		error_handler.RespondInternalErr(w, err)
		return
	}

	t := r.FormValue("type")
	file, header, err := r.FormFile("file")
	if err != nil {
		error_handler.RespondInternalErr(w, err)
		return
	}
	defer func() {
		_ = file.Close()
	}()

	tempFile, err := os.CreateTemp("static", fmt.Sprintf("upload-*%s", filepath.Ext(header.Filename)))
	if err != nil {
		error_handler.RespondInternalErr(w, err)
		return
	}

	f, err := h.usecase.Do(r.Context(), tempFile.Name(), t)
	if err != nil {
		error_handler.RespondInternalErr(w, err)
		return
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		error_handler.RespondInternalErr(w, err)
		return
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		error_handler.RespondInternalErr(w, err)
	}

	err = json.NewEncoder(w).Encode(f)
	if err != nil {
		error_handler.RespondInternalErr(w, err)
	}
}
