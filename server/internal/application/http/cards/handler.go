package cards

import (
	"encoding/json"
	"net/http"
	"strconv"

	error_handler "github.com/PakSerVal/bear-masha/internal/application/http"
	"github.com/PakSerVal/bear-masha/internal/application/http/cards/forms"
	"github.com/PakSerVal/bear-masha/internal/domain/cards/usecase/create"
	"github.com/PakSerVal/bear-masha/internal/domain/cards/usecase/delete"
	"github.com/PakSerVal/bear-masha/internal/domain/cards/usecase/get_by_id"
	"github.com/PakSerVal/bear-masha/internal/domain/cards/usecase/get_list"
	"github.com/PakSerVal/bear-masha/internal/domain/cards/usecase/update"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type CardHandler struct {
	getListUc get_list.Usecase
	getByIdUC get_by_id.Usecase
	createUC  create.Usecase
	deleteUC  delete.Usecase
	updateUC  update.Usecase
}

func New(
	getListUC get_list.Usecase,
	getByIdUC get_by_id.Usecase,
	createUC create.Usecase,
	deleteUC delete.Usecase,
	updateUC update.Usecase,
) *CardHandler {
	return &CardHandler{
		getListUc: getListUC,
		getByIdUC: getByIdUC,
		createUC:  createUC,
		deleteUC:  deleteUC,
		updateUC:  updateUC,
	}
}

func (h *CardHandler) GetList(w http.ResponseWriter, r *http.Request) {
	res, err := h.getListUc.Do(r.Context())
	if err != nil {
		error_handler.RespondInternalErr(w, err)
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		error_handler.RespondInternalErr(w, err)
	}
}

func (h *CardHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		error_handler.RespondErr(w, "handler: invalid id", http.StatusBadRequest)
		return
	}

	res, err := h.getByIdUC.Do(r.Context(), int64(id))
	if err != nil {
		error_handler.RespondInternalErr(w, err)
		return
	}

	if res == nil {
		error_handler.RespondErr(w, "card not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		error_handler.RespondInternalErr(w, err)
		return
	}
}

func (h *CardHandler) Create(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	cardForm := forms.CreateCard{}
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	err := decoder.Decode(&cardForm, r.PostForm)
	if err != nil {
		error_handler.RespondErr(w, "invalid params", http.StatusBadRequest)
		return
	}

	err = h.createUC.Do(
		r.Context(),
		cardForm.Title,
		cardForm.Description,
		cardForm.DeadlineAt,
		cardForm.Files,
	)
	if err != nil {
		error_handler.RespondInternalErr(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *CardHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		error_handler.RespondErr(w, "invalid card id", http.StatusBadRequest)
		return
	}

	err = h.deleteUC.Do(
		r.Context(),
		int64(id),
	)
	if err != nil {
		error_handler.RespondInternalErr(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *CardHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		error_handler.RespondErr(w, "invalid card id", http.StatusBadRequest)
		return
	}

	_ = r.ParseForm()
	updateForm := forms.UpdateCard{}
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	err = decoder.Decode(&updateForm, r.PostForm)
	if err != nil {
		error_handler.RespondErr(w, "invalid params", http.StatusBadRequest)
		return
	}

	err = h.updateUC.Do(
		r.Context(),
		int64(id),
		updateForm.Title,
		updateForm.Status,
		updateForm.Description,
		updateForm.DeadlineAt,
	)
	if err != nil {
		error_handler.RespondInternalErr(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
