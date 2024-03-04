package api

import (
	"awesome-backend-golang/internal/adapter/api/docs"
	"awesome-backend-golang/internal/aggregate/rooms"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"os"
)

type API struct {
	validate *validator.Validate
	rooms    rooms.RoomService
	//i18nSource *i18n.I18n
}

func New(opts ...func(*API)) *API {
	r := &API{}
	for _, opt := range opts {
		opt(r)
	}
	if r.validate == nil {
		r.validate = validator.New()
	}
	return r
}

func WithRooms(s rooms.RoomService) func(*API) {
	return func(r *API) { r.rooms = s }
}

func (a *API) NewRouter() chi.Router {
	info := docs.SwaggerInfo
	slog.Debug("api info", "title", info.Title, "version", info.Version, "path", info.BasePath)

	mux := chi.NewRouter()
	mux.Route("/room", func(r chi.Router) {
		r.Post("/getList", a.GetRoomList)
		r.Post("/register", a.RegisterRoom)
		r.Post("/modify", a.ModifyRoom)
		r.Post("/remove", a.RemoveRoom)
	})
	return mux
}

type ErrorResponse struct {
	Result  bool   `json:"result" xml:"result" form:"result"`
	Message string `json:"message" xml:"message" form:"message"`
}

func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	res := &ErrorResponse{Result: false, Message: err.Error()}
	var status int
	switch {
	case errors.Is(err, os.ErrInvalid):
		status = http.StatusBadRequest
	case errors.Is(err, os.ErrExist):
		status = http.StatusConflict
	case errors.Is(err, os.ErrNotExist):
		status = http.StatusNotFound
	default:
		status = http.StatusInternalServerError
		slog.ErrorContext(r.Context(), "unexpected", "error", err.Error())
	}
	render.Status(r, status)
	render.Respond(w, r, res)
}
