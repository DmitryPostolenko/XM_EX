package jwt

import (
	"os"
	"testing"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

var _ = os.Setenv("JWT_ACCESS_SECRET", "453RFW3cq$F3f#$6$%^$352V#$$$$@$Rc4WRFrWEF")

func TestCreateToken(t *testing.T) {
	userUuid, err := uuid.NewV4()
	if err != nil {
		t.Fatal("Could not create uuid, package: github.com/nu7hatch/gouuid")
	}

	tokenDetails, err := CreateToken(userUuid.String())
	if err != nil {
		t.Fatal("Could not create token, package: github.com/DmitryPostolenko/lets-go-chat/pkg/jwt")
	}

	tokenExpiresTime := time.Now().Add(time.Minute * 15).Unix()
	if tokenDetails.AccessUuid == "" || tokenDetails.AccessToken == "" || tokenDetails.Expires < tokenExpiresTime {
		t.Fatal("Wrong token details")
	}
}

func TestExtractTokenMetadataErr(t *testing.T) {
	tests := []struct {
		name     string
		token    string
		expected string
	}{
		{
			name:     "expired",
			token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjFkZTVkZDY2LTUwYmItNGI5Zi04MTMxLWU1NjVlMmU1NDA5ZiIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTYzODUyMDExNiwidXNlcl9pZCI6ImQzYWQ0MjFlLWU3OGMtNGI5Zi03YTUzLWJiMTA0MzM0MDkzMSJ9.oJFDDla_GQS0XET7aYNiYDnOvlRm6akR6BQa-5G_Xfw",
			expected: "Token is expired",
		},
		{
			name:     "invalid signature",
			token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjFkZTVkZDY2LTUwYmItNGI5Zi04MTMxLWU1NjVlMmU1NDA5ZiIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTYzODUyMDExNiwidXNlcl9pZCI6ImQzYWQ0MjFlLWU3OGMtNGI5Zi03YTUzLWJiMTA0Mzm0MDkzMSJ9.oJFDDla_GQS0XET7aYNiYDnOvlRm6akR6BQa-5G_Xfw",
			expected: "signature is invalid",
		},
		{
			name:     "invalid signing method",
			token:    "eyJHbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjFkZTVkZDY2LTUwYmItNGI5Zi04MTMxLWU1NjVlMmU1NDA5ZiIsImF1dGhvcml6ZWQiOnRydWUsImV4cCI6MTYzODUyMDExNiwidXNlcl9pZCI6ImQzYWQ0MjFlLWU3OGMtNGI5Zi03YTUzLWJiMTA0Mzm0MDkzMSJ9.oJFDDla_GQS0XET7aYNiYDnOvlRm6akR6BQa-5G_Xfw",
			expected: "signing method (alg) is unspecified.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := ExtractTokenMetadata(tt.token)
			if err == nil || err.Error() != tt.expected {
				t.Fatalf("Wrong error message! Expected: %v, actual: %v", tt.expected, err)
			}
		})
	}
}

func TestExtractTokenMetadataRes(t *testing.T) {
	userUuid := "d3ad421e-e78c-4b9f-7a53-bb1043340931"

	token, _ := CreateToken(userUuid)

	accessDetails, err := ExtractTokenMetadata(token.AccessToken)
	if err != nil || accessDetails.AccessUuid == "" || accessDetails.UserId != userUuid {
		t.Fatalf("Wrong accessDetails! %v", err)
	}
}
