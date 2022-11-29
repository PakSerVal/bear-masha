package http

import (
	"encoding/json"
	"net/http"

	"github.com/PakSerVal/bear-masha/internal/infrastucture/logger"
	"go.uber.org/zap"
)

func RespondErr(w http.ResponseWriter, errMsg string, code int) {
	errObj := struct {
		Error string `json:"error"`
	}{Error: errMsg}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(errObj)
	if err != nil {
		logger.Error("encoding response error", zap.Error(err))
	}
}

func RespondInternalErr(w http.ResponseWriter, err error) {
	logger.Error(err.Error(), zap.Error(err))

	RespondErr(w, "internal error", http.StatusInternalServerError)
}
