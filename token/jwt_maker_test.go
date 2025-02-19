package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/db/util"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	paylod, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, paylod)

	require.Equal(t, username, paylod.Username)
	require.WithinDuration(t, issuedAt, paylod.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, paylod.ExpiredAt, time.Second)
}

func TestJWTMakerExpiredToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.Equal(t, err.Error(), ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestJWTMakerInvalidTokenAlg(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomOwner(), time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	// tamper with the token
	token = token + "invalid"

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.Equal(t, err.Error(), ErrInvalidToken.Error())
	require.Nil(t, payload)
}
