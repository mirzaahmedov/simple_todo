package router

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	"github.com/mirzaahmedov/simple_todo/internal/store/mock"
)

func MakeTestHTTPRouter() *HTTPRouter {
	slog.SetDefault(
		slog.New(
			slog.NewTextHandler(os.Stdin, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}),
		),
	)

	r := NewHTTPRouter(
		mock.NewStore(),
		slog.Default(),
	)

	r.Init()

	return r
}

func MakeTestRequest(method string, url string, body any) (*http.Request, error) {
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}

		return http.NewRequest(method, url, bytes.NewReader(bodyBytes))
	}
	return http.NewRequest(method, url, nil)
}

func ParseJSON[T Model](buf *bytes.Buffer) (*SuccessResponse[T], error) {
	data := SuccessResponse[T]{}

	err := json.NewDecoder(buf).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
