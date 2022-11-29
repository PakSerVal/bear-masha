package auth

import (
	"encoding/json"
	"net/http"

	error_handler "github.com/PakSerVal/bear-masha/internal/application/http"
	"github.com/PakSerVal/bear-masha/internal/application/http/auth/forms"
	"github.com/PakSerVal/bear-masha/internal/domain/auth/usecase/sign_in"
	"github.com/PakSerVal/bear-masha/internal/infrastucture/logger"
	"github.com/gorilla/schema"
	"go.uber.org/zap"
)

type Handler struct {
	signInUC sign_in.Usecase
}

func New(signInUC sign_in.Usecase) *Handler {
	return &Handler{
		signInUC: signInUC,
	}
}

func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	authForm := forms.Authorize{}
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	err := decoder.Decode(&authForm, r.PostForm)
	if err != nil {
		logger.Error("invalid params", zap.Error(err))
		error_handler.RespondErr(w, "invalid params", http.StatusBadRequest)

		return
	}

	token, err := h.signInUC.Do(r.Context(), authForm.Login, authForm.Password)
	if err != nil {
		logger.Error("sign in usecase error", zap.Error(err))
		error_handler.RespondErr(w, "not authorized", http.StatusUnauthorized)

		return
	}

	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		error_handler.RespondInternalErr(w, err)
	}
}
