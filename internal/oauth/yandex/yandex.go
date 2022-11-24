package yandex

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/xopoww/wishes/internal/service"
)

type provider struct {
	c http.Client
	t Trace
	clientID string
}

func NewOAuthProvider(t Trace, clientID string) service.OAuthProvider {
	return &provider{
		t: t,
		clientID: clientID,
	}
}

type (
	OnValidateStartInfo struct {
		Req *http.Request
	}
	OnValidateDoneInfo struct {
		Resp  InfoResponse
		Error error
	}
)

func (p *provider) Validate(ctx context.Context, token string) (eid string, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://login.yandex.ru/info?format=json", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", fmt.Sprintf("OAuth %s", token))

	onDone := traceOnValidate(p.t, req.Clone(context.TODO()))
	var ir InfoResponse
	defer func() { onDone(ir, err) }()

	resp, err := p.c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("read body: %w", err)
		return
	}
	
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("status=%d, body=%q", resp.StatusCode, data)
		return
	}

	err = json.Unmarshal(data, &ir)
	if err != nil {
		err = fmt.Errorf("unmarshall %q: %w", data, err)
		return
	}

	if ir.ClientID != p.clientID {
		err = fmt.Errorf("client_id: got %q", ir.ClientID)
		return
	}
	
	return ir.ID, nil
}

type InfoResponse struct {
	Login    string `json:"login"`
	ID       string `json:"id"`
	ClientID string `json:"client_id"`
}

//go:generate gtrace

//gtrace:gen
//gtrace:set shortcut
type Trace struct {
	OnValidate func(OnValidateStartInfo) func(OnValidateDoneInfo)
}