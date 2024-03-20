package list

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"url-shortener/internal/lib/api/response"
	"url-shortener/internal/lib/sl"
	"url-shortener/internal/storage"
)

func GetAll(log *slog.Logger, urlLister storage.URLLister) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.list.GetAll"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		urls, err := urlLister.GetAllURLs()
		if err != nil {
			log.Error("Failed to get URLs", sl.Err(err))
			render.JSON(w, r, response.Error("Failed to get URLs"))
			return
		}

		render.JSON(w, r, urls)
	}
}
