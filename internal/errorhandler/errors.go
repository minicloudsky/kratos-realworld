package errorhandler

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	nethttp "net/http"
)

func NewHTTPError(code int, field string, detail string) *HTTPError {
	return &HTTPError{
		Errors: map[string][]string{
			field: {detail},
		},
		Code: code,
	}
}

type HTTPError struct {
	Errors map[string][]string `json:"errorhandler"`
	Code   int                 `json:"code"`
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTPError %d", e.Code)
}

func FromError(err error) *HTTPError {
	if err == nil {
		return nil
	}
	if se := new(HTTPError); errors.As(err, &se) {
		return se
	}

	return &HTTPError{}
}

func ErrorEncoder(w nethttp.ResponseWriter, r *nethttp.Request, err error) {
	se := FromError(err)
	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/"+codec.Name())
	w.WriteHeader(se.Code)
	_, _ = w.Write(body)
}
