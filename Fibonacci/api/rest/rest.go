package rest

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/rs/zerolog"

	"github.com/rauf95/rauf/core"
)

type api struct {
	logger zerolog.Logger
}

func New(logger zerolog.Logger) http.Handler {
	a := api{
		logger: logger,
	}

	mu := http.NewServeMux()
	mu.HandleFunc("/fibonacci", a.fibonacciHandler)

	return mu
}

func (a *api) fibonacciHandler(w http.ResponseWriter, r *http.Request) {
	arg, ok := r.URL.Query()["arg"]
	if !ok || len(arg) == 0 {
		a.logger.Warn().Msg("url param 'arg' is missing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	intArg, err := strconv.Atoi(strings.Join(arg, ""))
	if err != nil {
		a.logger.Warn().Strs("arg", arg).Msg("invalid arg")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fibonacci := core.Fibonacci(intArg)

	err = json.NewEncoder(w).Encode(fibonacci)
	if err != nil {
		a.logger.Error().Err(err).Msg("internal error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
