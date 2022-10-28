package auth

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Adeithe/go-twitch/api"
)

func Login(api *api.Client) (*api.TwitchLogin, error) {
	fmt.Print("Twitch Username: ")
	username := stdin()
	fmt.Print("Twitch Password: ")
	password := stdin()

	login, err := api.Login(username, password)
	if err != nil {
		return nil, fmt.Errorf("api authentication error: %w", err)
	}

	errCode := login.GetErrorCode()
	fmt.Printf("Tried to login, response code: %d\n", errCode)
	if errCode == 3011 || errCode == 3022 {
		fmt.Print("Twitch 2FA: ")
		code := stdin()
		if err := login.Verify(code); err != nil {
			return nil, fmt.Errorf("%d, %w", errCode, err)
		}
	} else if errCode != 0 {
		//nolint:goerr113
		return nil, fmt.Errorf("failed: %s", login.GetError())
	}

	fmt.Printf("Twitch Access Token: %s\n", login.GetAccessToken())

	return login, nil
}

func stdin() string {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')

	return strings.TrimSuffix(strings.TrimSuffix(str, "\r\n"), "\n")
}
