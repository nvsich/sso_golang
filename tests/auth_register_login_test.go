package tests

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang-jwt/jwt/v5"
	ssov1 "github.com/nvsich/sso_protos/gen/go/sso"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sso/tests/suite"
	"testing"
	"time"
)

const (
	emptyAppId = 0
	appID      = 1
	appSecret  = "test secret"

	passDefaultLen = 10

	deltaSeconds = 1
)

func TestRegisterLogin_Login_HappyPath(t *testing.T) {
	ctx, s := suite.New(t)

	tEmail := gofakeit.Email()
	tPassword := randomPassword()

	// Register
	regResponse, err := s.AuthClient.Register(ctx, &ssov1.RegisterRequest{
		Email:    tEmail,
		Password: tPassword,
	})

	require.NoError(t, err)
	require.NotEmpty(t, regResponse.GetUserId())

	// Login
	loginResponse, err := s.AuthClient.Login(ctx, &ssov1.LoginRequest{
		Email:    tEmail,
		Password: tPassword,
	})

	loginTime := time.Now()

	require.NoError(t, err)

	token := loginResponse.GetToken()
	require.NotEmpty(t, token)

	tokenParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(appSecret), nil
	})

	require.NoError(t, err)

	claims, ok := tokenParsed.Claims.(jwt.MapClaims)

	assert.True(t, ok)
	assert.Equal(t, regResponse.GetUserId(), int64(claims["uid"].(float64)))
	assert.Equal(t, tEmail, claims["email"].(string))
	assert.Equal(t, appID, claims["app_id"].(float64))
	assert.InDelta(t, loginTime.Add(s.Cfg.TokenTTL).Unix(), claims["exp"].(float64), deltaSeconds)
}

func randomPassword() string {
	return gofakeit.Password(true, true, true, true, true, passDefaultLen)
}
