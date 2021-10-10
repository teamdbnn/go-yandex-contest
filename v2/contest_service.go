package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/teamdbnn/go-yandex-contest/v2/common"
)

// GetContestItemService Get Contest item info
type GetContestItemService struct {
	c       *Client
	contest int64
}

// Contest Set contest id
func (s *GetContestItemService) Contest(contest int64) *GetContestItemService {
	s.contest = contest
	return s
}

// Do Send request
func (s *GetContestItemService) Do(ctx context.Context, opts ...RequestOption) (res *ContestDescription, err error) {
	if s.contest == 0 {
		return nil, common.ValidationRequiredError("Contest")
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%d", s.contest),
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
