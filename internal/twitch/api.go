package twitch

import (
	"context"
	"fmt"

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
		return nil, fmt.Errorf("failed to retrieve oauth2 token: %w", err)
	}

	return api.NewBearer(token.AccessToken), nil
}
