package authcode

import (
	"context"
	"fmt"
	"net/http"
)

func WaitForAuthorizationCode() (string, error) {

	m := http.NewServeMux()
	s := &http.Server{Addr: ":42001", Handler: m}

	var code = ""

	m.HandleFunc("/exchange_token", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, "The authorization code has been received. You can now close this window.")

		code = r.URL.Query().Get("code")

		go func() {
			s.Shutdown(context.Background())
		}()
	})

	s.ListenAndServe()

	return code, nil
}
