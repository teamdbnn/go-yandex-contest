package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type UserGeneratePasswordService struct {
	c    *Client
	user int64
}

func (s *UserGeneratePasswordService) User(user int64) *UserGeneratePasswordService {
	s.user = user
	return s
}

// Do Send request
func (s *UserGeneratePasswordService) Do(ctx context.Context, opts ...RequestOption) (res *ContestDescription, err error) {
	if s.user == 0 {
		return nil, errors.New("Set user(%d)")
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("/user/%d/generate-password", s.user),
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ContestDescription)

	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
