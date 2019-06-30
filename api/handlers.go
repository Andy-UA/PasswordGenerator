package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../domain"
	"../service"
)

func getPasswordConfigs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//check on error defer r.Body.Close() (https://golang.org/pkg/net/http/#Request - The Server will close the request body. The ServeHTTP Handler does not need to.)

	if configs := r.Context().Value("passwordConfig"); configs != nil {

		response := service.GeneratePassword(configs.(domain.PasswordConfig))

		responseInstance := domain.Response{
			GeneratedPassword: response,
		}
		resp, err := json.Marshal(responseInstance)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Context is empty", http.StatusInternalServerError)
		return
	}
}

func httpMethodHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid method", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func bodyAvailabilityHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength <= 0 {
			http.Error(w, "request body is empty", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func bodyContentHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var passwordConfigs domain.PasswordConfig
		if err = json.Unmarshal(body, &passwordConfigs); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err = CheckBodyContent(passwordConfigs); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "passwordConfig", passwordConfigs)))
	})
}

func CheckBodyContent(configs domain.PasswordConfig) error {
	if configs.MinLength < 0 || configs.NumberAmount < 0 || configs.SpecialCharsAmount < 0 {
		return fmt.Errorf("parameter's value can't be less than 0")
	}
	return nil
}
