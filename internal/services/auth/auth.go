package auth

import (
	"SPRk_Space/internal/config"
	"SPRk_Space/internal/entities"
	"SPRk_Space/internal/logger"
	"SPRk_Space/internal/storage"
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/argon2"
)

type Auth struct {
	ctx    context.Context
	config *config.Auth
	logger logger.Logger

	pepper []byte
	st     *storage.Storage
}

func newAuth(ctx context.Context, cfg *config.Auth, log logger.Logger, st *storage.Storage) (*Auth, error) {
	return &Auth{
		ctx:    ctx,
		config: cfg,
		logger: log,
		st:     st,
		pepper: nil,
	}, nil
}

func (a *Auth) Login(ctx context.Context, login string, pass string) (string, error) {
	key := deriveAuthKey(login, pass, a.pepper)

	userId, err := a.st.Login.FindUserByAuthKey(key)
	if err != nil {
		return "", err
	}

	token, tokenHash, err := newToken()
	if err != nil {
		return "", err
	}

	session := entities.Session{
		Token:       tokenHash,
		UserId:      userId,
		CreateTime:  time.Now(),
		TTL:         0,
		Fingerprint: entities.Fingerprint{},
	}
	err = a.st.Session.SaveSession(a.ctx, session)
	if err != nil {
		return "", err
	}
	return token, nil
}

func deriveAuthKey(login, passwordHash string, pepper []byte) string {
	login = strings.ToLower(strings.TrimSpace(login))

	m := hmac.New(sha256.New, pepper)
	m.Write([]byte(login))
	salt := m.Sum(nil)

	key := argon2.IDKey([]byte(passwordHash), salt, 3, 64*1024, 2, 32)
	return hex.EncodeToString(key)
}

func newToken() (token string, tokenHash string, err error) {
	b := make([]byte, 32)
	_, err = rand.Read(b)
	if err != nil {
		return "", "", fmt.Errorf("Auth: NewToken: Rand: %w", err)
	}
	token = base64.RawURLEncoding.EncodeToString(b)
	sum := sha256.Sum256([]byte(token))
	tokenHash = hex.EncodeToString(sum[:])
	return token, tokenHash, nil
}
