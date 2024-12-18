package main

import (
	"log/slog"
	"net/http"
)

func logRequest(logger *slog.Logger, r *http.Request) {
	logger.Info("request made", slog.String("url", r.URL.String()),
		slog.String("proto", r.Proto),
		slog.String("method", r.Method),
		slog.String("remote-addr", r.RemoteAddr))
}

// https://owasp.org/www-project-secure-headers/ci/headers_add.json
func (app *application) secureHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Content-Security-Policy", `default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com`)
		// w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")

		h.ServeHTTP(w, r)
	})
}

// func (app *application) excerptRequired(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		params := httprouter.ParamsFromContext(r.Context())
// 		idStr := params.ByName("excerpt")
// 		id, err := strconv.ParseInt(idStr, 10, 64)
// 		if err != nil {
// 			app.clientError(w, http.StatusBadRequest)
// 			return
// 		}
// 		if int(id) >= len(app.excerpts) {
// 			app.clientError(w, http.StatusBadRequest)
// 			return
// 		}

// 		ctx := context.WithValue(r.Context(), excerptKey, app.excerpts[id])
// 		ctx = context.WithValue(ctx, excerptIdKey, int(id))
// 		h.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }
