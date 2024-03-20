package delete

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"url-shortener/internal/lib/sl"

	"github.com/go-chi/render"
	"log/slog"
	"url-shortener/internal/lib/api/response"
	"url-shortener/internal/storage"
)

func DeleteURL(log *slog.Logger, urlDeleter storage.URLDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.delete.DeleteURL"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")

		err := urlDeleter.DeleteURL(alias)
		if err != nil {
			log.Error("Failed to delete URL", sl.Err(err))
			render.JSON(w, r, response.Error("Failed to delete URL"))
			return
		}

		render.JSON(w, r, response.OK())
	}
}
