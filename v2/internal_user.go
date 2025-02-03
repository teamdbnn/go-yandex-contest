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
		return requiredFieldError("userID")
	}
	return nil
}

// Do Generate new password for internal user
//
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/internal-user/generateInternalUserPasswordUsingPOST
// meta:operation POST /user/{userId}/generate-password
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
	if err = json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}
