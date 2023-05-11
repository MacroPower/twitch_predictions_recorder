package twitch

import (
	"context"

	"github.com/Adeithe/go-twitch"
	"github.com/Adeithe/go-twitch/api"
	"golang.org/x/oauth2/clientcredentials"
	oauth2 "golang.org/x/oauth2/twitch"
)

func NewAPIClient(clientID, clientSecret string) (*api.Client, error) {
	api := twitch.API(clientID)

	oauth2Config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     oauth2.Endpoint.TokenURL,
	}

	token, err := oauth2Config.Token(context.Background())
	if err != nil {
		return nil, err
	}

	return api.NewBearer(token.AccessToken), nil
}
