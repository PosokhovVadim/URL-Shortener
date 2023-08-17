package save

import (
	"fmt"
	"net/http"
	resp "url-shortener/internal/http-server/handlers/response"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/lib/random"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"golang.org/x/exp/slog"
)

const aliasLength = 4

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}

type URLSaver interface {
	InsertOneURL(url, alias string) error
	InsertManyURL(values map[string]string) error
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "response.url.save.New"

		log = slog.With(
			slog.String("func", fn),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request
		err := render.DecodeJSON(r.Body, &req)

		if err != nil {
			log.Error(fmt.Sprintf("DecodeJSON error in func: %s", fn), sl.Err(err))

			render.JSON(w, r, resp.Error("DecodeJSON error"))
			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		//Validate
		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)
			log.Error("Error validating request", sl.Err(err))

			render.JSON(w, r, resp.ValidationError(validateErr))
			return

		}

		alias := req.Alias
		if alias == "" {
			alias = random.RandomString(aliasLength)
		}
	}

}
