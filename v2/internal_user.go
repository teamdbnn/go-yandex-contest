package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// UserGeneratePassword Generate user password
type UserGeneratePassword struct {
	c      *Client
	userID int64
}

// UserID Set user id
func (s *UserGeneratePassword) UserID(userID int64) *UserGeneratePassword {
	s.userID = userID
	return s
}

func (s *UserGeneratePassword) validate() error {
	if s.userID == 0 {
		return requiredError("userID")
	}
	return nil
}

// Do Send POST request
func (s *UserGeneratePassword) Do(ctx context.Context, opts ...RequestOption) (*UserWithPasswordResponse, error) {
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
