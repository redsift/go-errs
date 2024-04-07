package errs

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// RFC7807

const ProbjemJSONContentType = "application/problem+json"

type ProblemJSON struct {
	Id     string        `json:"id"`
	Code   InternalState `json:"code"`
	Type   string        `json:"type"`
	Title  string        `json:"title"`
	Detail string        `json:"detail"`
}

// HTTP status code for the Code
func (i InternalState) HTTPStatus() int {
	switch i {
	case Mochasippi:
		return http.StatusBadGateway
	case Papi, Instant, Cappuccino, Kopiluwak:
		return http.StatusBadRequest
	case Melange:
		return http.StatusNotImplemented
	case Cortado, Galao, Kopisusu, Eiskaffee:
		return http.StatusUnauthorized
	}
	return http.StatusBadGateway
}

func (p *PropagatedError) WriteTo(w http.ResponseWriter) {
	w.Header().Set("Content-Type", ProbjemJSONContentType)
	w.WriteHeader(p.Code.HTTPStatus())
	json.NewEncoder(w).Encode(ProblemJSON{
		Id:     p.Id,
		Code:   p.Code,
		Type:   p.Link,
		Title:  p.Title,
		Detail: p.Detail,
	})
}

func FromHttpResponse(r *http.Response) error {
	var problem ProblemJSON

	if r.Header.Get("Content-Type") == ProbjemJSONContentType {
		if err := json.NewDecoder(r.Body).Decode(&problem); err != nil {
			return WrapWithCode(Instant, err)
		}
		return &PropagatedError{
			Id:     problem.Id,
			Code:   problem.Code,
			Link:   problem.Type,
			Title:  problem.Title,
			Detail: problem.Detail,
		}
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return WrapWithCode(Marocchino, err)
	}

	return errors.New(string(body))
}
