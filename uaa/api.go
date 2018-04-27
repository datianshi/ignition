package uaa

import (
	"context"

	"github.com/cloudfoundry-incubator/uaa-cli/uaa"
)

// API is used to access a UAA server
type API interface {
	UserIDForAccountName(a string) (string, error)
	CreateUser(username, origin, externalID, email string) (string, error)
}

// Authenticate will authenticate with a UAA server and set the Token and Client
// for the UAAAPI
func (a *Client) Authenticate() error {
	if a.Client == nil {
		a.Client = a.OauthConfig.Client(context.Background())
	}

	token, err := a.OauthConfig.Token(context.Background())
	if err != nil {
		return err
	}
	if a.userManager == nil {
		uaaConfig := uaa.NewConfig()
		uaaConfig.AddTarget(uaa.Target{BaseUrl: a.URL})
		uaaConfig.AddContext(uaa.NewContextWithToken(token.AccessToken))
		a.userManager = &uaa.UserManager{
			Config:     uaaConfig,
			HttpClient: a.Client,
		}
	}
	return nil
}
