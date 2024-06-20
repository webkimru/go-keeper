package jwtmanager

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJWTManager_BuildJWTString(t *testing.T) {
	tests := []struct {
		name      string
		secretKey string
		tokenExp  int
		userID    int64
		wantErr   bool
	}{
		{"positive: correct data", "secret", 1, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := JWTManager{
				secretKey: tt.secretKey,
				tokenExp:  tt.tokenExp,
			}
			_, err := j.BuildJWTString(tt.userID)
			if (err != nil) != tt.wantErr {
				assert.NoError(t, err)
			}
		})
	}
}

func TestJWTManager_GetUserID(t *testing.T) {
	j := JWTManager{
		secretKey: "secret",
		tokenExp:  1,
	}
	token, _ := j.BuildJWTString(1)

	tests := []struct {
		name    string
		secret  string
		token   string
		want    int64
		wantErr bool
	}{
		{
			"positive: correct data",
			"secret",
			token,
			1,
			false,
		},
		{
			"negative: expired token",
			"secret",
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTg4MDg0MTAsIlVzZXJJRCI6MX0.9GfhPnfWiifquQFE0kvWRQsOinqSh5n2ajHp1RIaOP4",
			-1,
			true,
		},
		{
			"negative: token alg none",
			"secret",
			"eyJhbGciOiJub25lIn0.eyJleHAiOjE3MTg4MDg0MTAsIlVzZXJJRCI6MX0.9GfhPnfWiifquQFE0kvWRQsOinqSh5n2ajHp1RIaOP4",
			-1,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := JWTManager{secretKey: tt.secret}
			assert.Equalf(t, tt.want, j.GetUserID(tt.token), "GetUserID(%v)", tt.token)
		})
	}
}
