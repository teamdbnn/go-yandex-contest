package yacontest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/teamdbnn/go-yandex-contest/v2/common"
)

type SubmissionRejudgeService struct {
	c          *Client
	submission int64
}

func (s *SubmissionRejudgeService) Submission(submission int64) *SubmissionRejudgeService {
	s.submission = submission
	return s
}

func (s *SubmissionRejudgeService) Do(ctx context.Context, opts ...RequestOption) (*Empty, error) {
	if s.submission == 0 {
		return nil, common.ValidationRequiredError("Sumbission")
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("/submissions/%v/rejudge", s.submission),
	}
	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return new(Empty), nil
}
