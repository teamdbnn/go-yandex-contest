package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserGeneratePasswordService struct {
	c      *Client
	userID int64
}

func (s *UserGeneratePasswordService) UserID(userID int64) *UserGeneratePasswordService {
	s.userID = userID
	return s
}

func (s *UserGeneratePasswordService) validate() error {
	if s.userID == 0 {
		return requiredError("userID")
	}
	return nil
}

// Do Send request
func (s *UserGeneratePasswordService) Do(ctx context.Context, opts ...RequestOption) (*UserWithPasswordResponse, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("/user/%d/generate-password", s.userID),
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(UserWithPasswordResponse)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
