package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/teamdbnn/go-yandex-contest/v2/models"
)

type GetContestItemService struct {
	c       *Client
	contest int64
}

func (s *GetContestItemService) Contest(contest int64) *GetContestItemService {
	s.contest = contest
	return s
}

func (s *GetContestItemService) Do(ctx context.Context, opts ...RequestOption) (res *models.Contest, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%d", s.contest),
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(models.Contest)

	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
