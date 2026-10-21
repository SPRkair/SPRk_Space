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
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/crypto/argon2"
)

var (
	ErrInvalidInput  = errors.New("Invalid Input")
	ErrTokenNotFound = errors.New("Token not found")
)

type Auth struct {
	ctx    context.Context
	config *config.Auth
	logger logger.Logger

	pepper []byte

	session storage.SessionStorageDriver
	login   storage.LoginStorageDriver
	user    storage.UserStorageDriver
}

func NewAuth(ctx context.Context, cfg *config.Auth, log logger.Logger, st *storage.Storage) *Auth {
	return &Auth{
		ctx:     ctx,
		config:  cfg,
		logger:  log,
		session: st.Session,
		login:   st.Login,
		user:    st.User,
		pepper:  nil,
	}
}

func (a *Auth) RegstrationNewUser(ctx context.Context, ri entities.RegistrationInput) (string, error) {
	err := a.validateRegInput(ri)
	if err != nil {
		return "", ErrInvalidInput
	}

	key := deriveAuthKey(ri.Login, ri.Pass, a.pepper)

	userId, err := a.login.RegNewLogin(key)
	if err != nil {
		return "", err
	}

	user := entities.User{
		Id:           userId,
		Name:         ri.UserName,
		Email:        ri.Email,
		Status:       entities.Registered,
		Registration: time.Now(),
		Permissions:  make(map[string]entities.Permissions),
	}

	err = a.user.SaveUser(a.ctx, user)
	if err != nil {
		return "", err
	}
	var token, tokenHash string
	for {
		token, tokenHash, err = newToken()
		if err != nil {
			return "", err
		}
		_, err = a.session.Verify(a.ctx, tokenHash)
		if err == ErrTokenNotFound {
			break
		}
		if err != nil {
			return "", err
		}
	}

	session := entities.Session{
		Token:       tokenHash,
		UserId:      userId,
		CreateTime:  time.Now(),
		TTL:         0,
		Fingerprint: entities.Fingerprint{},
	}
	err = a.session.SaveSession(a.ctx, session)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *Auth) Login(ctx context.Context, li entities.LoginInput) (string, error) {
	key := deriveAuthKey(li.Login, li.Pass, a.pepper)

	userId, err := a.login.FindUserByAuthKey(key)
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
	err = a.session.SaveSession(a.ctx, session)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *Auth) validateRegInput(ri entities.RegistrationInput) error {
	login := strings.TrimSpace(ri.Login)
	lc := utf8.RuneCountInString(login)
	if lc < a.config.MinLoginRunes || lc > a.config.MaxLoginRunes {
		return ErrInvalidInput
	}
	if !utf8.ValidString(login) {
		return ErrInvalidInput
	}

	unc := utf8.RuneCountInString(ri.UserName)
	if unc < a.config.MinUserNameRunes || unc > a.config.MaxUserNameRunes {
		return ErrInvalidInput
	}
	if !utf8.ValidString(ri.UserName) {
		return ErrInvalidInput
	}

	if utf8.RuneCountInString(ri.Pass) < a.config.MinPassLen {
		return ErrInvalidInput
	}
	return nil
}

func (a *Auth) Close() {
	a.logger.Debug_info("Auth Close")
}

func deriveAuthKey(login string, passwordHash string, pepper []byte) string {
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
