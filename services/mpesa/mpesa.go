package mpesa

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

type MpesaService struct {
	ConsumerKey    string
	ConsumerSecret string
	EndPoint       string
}

func NewMpesaService(ConsumerKey, consumerSecret, endPoint string) *MpesaService {
	return &MpesaService{
		ConsumerKey:    ConsumerKey,
		ConsumerSecret: consumerSecret,
		EndPoint:       endPoint,
	}
}

type authResponse struct {
	AccessToken string `json:"access_token"`
}

func (m *MpesaService) auth() (string, error) {
	auth := base64.StdEncoding.EncodeToString([]byte(m.ConsumerKey + ":" + m.ConsumerSecret))
	url := m.EndPoint + "/oauth/v1/generate?grant_type=client_credentials"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("could not create request: %v", err)
	}
	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("could not send request: %v", err)
	}
	if resp.Body == nil {
		return "", fmt.Errorf("empty response")
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return "", fmt.Errorf("invalid request client secret or key")
	}
	// timeouts, network errors, etc
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("could not authenticate: %v", resp.Status)
	}

	var authResp authResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return "", fmt.Errorf("could not decode response: %v", err)
	}
	return authResp.AccessToken, nil
}




