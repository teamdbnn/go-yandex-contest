package yacontest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client Define API client
type Client struct {
	OAuthToken string
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
	Debug      bool
	Logger     logger
	TimeOffset int64
	do         doFunc
}
type doFunc func(req *http.Request) (*http.Response, error)
type logger func(format string, v ...interface{})

const (
	baseAPIMainURL = "https://api.contest.yandex.net/api/public/v2"
)

// NewClient Create new http client
func NewClient(outhToken string) *Client {
	return &Client{
		OAuthToken: outhToken,
		BaseURL:    baseAPIMainURL,
		UserAgent:  "Yandex-Contest/golang",
		HTTPClient: http.DefaultClient,
	}
}

// WithLogger add logger
func (c *Client) WithLogger(l logger) *Client {
	c.Logger = l
	return c
}

// debug Put data in log
func (c *Client) debug(format string, v ...interface{}) {
	if !c.Debug || c.Logger == nil {
		return
	}

	c.Logger(format, v...)
}

func (c *Client) parseRequest(r *request, opts ...RequestOption) (err error) {
	for _, opt := range opts {
		opt(r)
	}
	err = r.validate()
	if err != nil {
		return err
	}
	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)

	queryString := r.query.Encode()
	body := &bytes.Buffer{}
	bodyString := r.form.Encode()
	header := http.Header{}
	if r.header != nil {
		header = r.header.Clone()
	}

	if bodyString != "" {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		body = bytes.NewBufferString(bodyString)
	} else if r.json != nil {
		header.Set("Content-Type", "application/json")
		jsonString, _ := json.Marshal(r.json)
		body = bytes.NewBuffer(jsonString)
	}

	header.Set("Authorization", fmt.Sprintf("%s %s", "OAuth", c.OAuthToken))

	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	c.debug("full url: %s, body: %s", fullURL, bodyString)
	r.fullURL = fullURL
	r.header = header
	r.body = body
	return nil
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	c.debug("request: %#v", req)
	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}
	res, err := f(req)
	if err != nil {
		return []byte{}, err
	}
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		if err == nil && cerr != nil {
			err = cerr
		}
	}()
	c.debug("response: %#v", res)
	c.debug("response body: %#v", string(data))
	c.debug("response status code: %#v", res.StatusCode)

	if res.StatusCode >= http.StatusBadRequest {
		apiErr := &APIError{
			Code:    int64(res.StatusCode),
			Message: string(data),
		}
		return nil, apiErr
	}

	return data, nil
}

// NewGetContestService Init contest item Service
func (c *Client) NewGetContestService() *GetContestService {
	return &GetContestService{c: c}
}

// NewsGetParticipantListOfCompetition Init Participant List of Competition service
func (c *Client) NewGetParticipantsOfCompetitionService() *GetParticipantsOfCompetitionService {
	return &GetParticipantsOfCompetitionService{c: c}
}

// NewGetParticipantListOfContest Init Participant List of contest service
func (c *Client) NewGetParticipantsOfContestService() *GetParticipantsOfContestService {
	return &GetParticipantsOfContestService{c: c}
}

// NewUserGeneratePasswordService Init Generate user password service
func (c *Client) NewUserGeneratePasswordService() *UserGeneratePasswordService {
	return &UserGeneratePasswordService{c: c}
}

func (c *Client) NewSubmissionRejudgeService() *SubmissionRejudgeService {
	return &SubmissionRejudgeService{c: c}
}

func (c *Client) NewRegisterParticipantIntoCompetitionService() *RegisterParticipantIntoCompetitionService {
	return &RegisterParticipantIntoCompetitionService{c: c}
}

func (c *Client) NewGetClarificationsInContestService() *GetClarificationsInContestService {
	return &GetClarificationsInContestService{c: c}
}

func (c *Client) NewGetContestMessagesServie() *GetContestMessagesServie {
	return &GetContestMessagesServie{c: c}
}

func (c *Client) NewRegisterParticipantForContestService() *RegisterParticipantForContestService {
	return &RegisterParticipantForContestService{c: c}
}

func (c *Client) NewUpdateParticipantForContestService() *UpdateParticipantForContestService {
	return &UpdateParticipantForContestService{c: c}
}
