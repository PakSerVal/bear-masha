package middleware

import (
	"net/http"

	"github.com/PakSerVal/bear-masha/internal/domain/auth/usecase/validate"
)

type authenticationMiddleware struct {
	authUC validate.Usecase
}

func New(authUC validate.Usecase) *authenticationMiddleware {
	return &authenticationMiddleware{authUC: authUC}
}

func (amw *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() == "/auth" {
			next.ServeHTTP(w, r)
			return
		}

		token := r.Header.Get("X-Token")

		if err := amw.authUC.Do(r.Context(), token); err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
