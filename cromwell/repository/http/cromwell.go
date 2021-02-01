package http

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/lmtani/cromwell-cli/domain"
)

type httpCromwellRepository struct {
	c httpClient
}

func NewHTTPCromwellRepository(host, token string) domain.CromwellRepository {
	c := newHttpClient(host, token)
	return &httpCromwellRepository{c}
}

// func (h *httpCromwellRepository) Submit(ctx context.Context) domain.SubmitResponse {

// }

// func (h *httpCromwellRepository) Kill(ctx context.Context) domain.SubmitResponse {}

func (h *httpCromwellRepository) Query() (r domain.WorkflowQueryResponse, err error) {
	route := "/api/workflows/v1/query"
	params := url.Values{}
	params.Add("includeSubworkflows", "true")
	resp, err := h.c.Get(route + "?" + params.Encode())
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return r, err
	}

	if resp.StatusCode >= 400 {
		return r, fmt.Errorf("Submission failed. The server returned %d\n%#v", resp.StatusCode, r)
	}
	return
}
