package mocks

import (
	"fmt"

	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/config"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
)

type MockClientConfig struct {
	ProfileNameFunc    func() string
	EnvironmentFunc    func() string
	ClientIDFunc       func() string
	ClientSecretFunc   func() string
	RedirectURIFunc    func() string
	OAuthTokenDataFunc func() string
	AccessTokenFunc    func() string
	LogFilePathFunc    func() string
	LoggingEnabledFunc func() bool
}

var UpdatedAccessToken string

func (m *MockClientConfig) ProfileName() string {
	return m.ProfileNameFunc()
}

func (m *MockClientConfig) Environment() string {
	return m.EnvironmentFunc()
}

func (m *MockClientConfig) ClientID() string {
	return m.ClientIDFunc()
}

func (m *MockClientConfig) ClientSecret() string {
	return m.ClientSecretFunc()
}

func (m *MockClientConfig) RedirectURI() string {
	return m.RedirectURIFunc()
}

func (m *MockClientConfig) OAuthTokenData() string {
	return m.OAuthTokenDataFunc()
}

func (m *MockClientConfig) AccessToken() string {
	return m.AccessTokenFunc()
}

func (m *MockClientConfig) LogFilePath() string {
	return m.LogFilePathFunc()
}

func (m *MockClientConfig) LoggingEnabled() bool {
	return m.LoggingEnabledFunc()
}

func (m *MockClientConfig) String() string {
	return fmt.Sprintf("\n-------------\nProfile Name: %s\nEnvironment: %s\nLogging Enabled: %v\nLog File Path: %s\nClient ID: %s\nClient Secret: %s\nRedirect URI: %s\nAccess Token: %s\n--------------\n", m.ProfileName(), m.Environment(), m.LoggingEnabled(), m.LogFilePath(), m.ClientID(), m.ClientSecret(), m.RedirectURI(), m.AccessToken())
}

func UpdateOAuthToken(_ config.Configuration, oauthTokenData *models.OAuthTokenData) error {
	UpdatedAccessToken = oauthTokenData.AccessToken
	return nil
}

func OverridesApplied() bool {
	return false
}
