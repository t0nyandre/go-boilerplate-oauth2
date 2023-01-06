package github

import (
	"context"
	"log"

	"github.com/t0nyandre/go-boilerplate-oauth2/pkg/session"
	"github.com/t0nyandre/go-boilerplate-oauth2/pkg/utils/state"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type Provider struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	CallbackURL  string `json:"callback_url"`
	State        string
	AuthURL      string
	TokenURL     string
	Config       *oauth2.Config
}

func New(ctx context.Context, clientID, clientSecret, callbackURL string, scopes ...string) *Provider {
	logger, ok := ctx.Value("logger").(*zap.SugaredLogger)
	if !ok {
		log.Fatalln("Could not get logger from context")
	}

	state, err := state.GenerateRandomState()
	if err != nil {
		logger.Fatalw("Could not generate random state", "error", err)
	}

	provider := &Provider{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		CallbackURL:  callbackURL,
		State:        state,
		AuthURL:      "https://github.com/login/oauth/authorize",
		TokenURL:     "https://github.com/login/oauth/access_token",
	}
	provider.Config = provider.newConfig(scopes)

	return provider
}

func (p *Provider) newConfig(scopes []string) *oauth2.Config {
	c := &oauth2.Config{
		ClientID:     p.ClientID,
		ClientSecret: p.ClientSecret,
		RedirectURL:  p.CallbackURL,
		Scopes:       []string{},
		Endpoint: oauth2.Endpoint{
			AuthURL:  p.AuthURL,
			TokenURL: p.TokenURL,
		},
	}

	for _, scope := range scopes {
		c.Scopes = append(c.Scopes, scope)
	}

	return c
}

func (p *Provider) getAccessToken(code string) (*session.TokenData, error) {
	token, err := p.Config.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	tokenData := &session.TokenData{
		Token:   token.AccessToken,
		Expires: token.Expiry,
	}

	return tokenData, nil
}
