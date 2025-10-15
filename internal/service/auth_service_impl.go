package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"it-backend/internal/helper"
	"it-backend/internal/model/dto"
	"it-backend/internal/repository"
	"time"

	"github.com/go-ldap/ldap/v3"
	"github.com/pkg/errors"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
}

func NewAuthService(authRepository repository.AuthRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
	}
}

func (s *AuthServiceImpl) Login(ctx context.Context, request dto.LoginRequest) (string, error) {
	user, err := s.AuthRepository.FindUserByUsername(ctx, request.Username)
	if err != nil {
		return "", errors.Wrap(err, "invalid credentials")
	}

	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
	// 	return "", errors.Wrap(err, "invalid credentials")
	// }

	accessToken, err := helper.GenerateToken(ctx, user, "access_token", time.Now())
	if err != nil {
		return "", errors.Wrap(err, "failed to generate token")
	}

	return accessToken, nil
}

func (s *AuthServiceImpl) LoginLDAP(ctx context.Context, request dto.LoginRequest) (string, error) {
	cfg, err := helper.GetConfig()
	if err != nil {
		return "", errors.Wrap(err, "gagal load config")
	}

	ldapURL := fmt.Sprintf("%s:%d", cfg.LDAP_SERVER, cfg.LDAP_PORT)
	dn := fmt.Sprintf("uid=%s,%s,%s", ldap.EscapeFilter(request.Username), cfg.LDAP_USER_DN, cfg.LDAP_BASE_DN)

	conn, err := ldap.DialURL(ldapURL)
	if err != nil {
		return "", errors.Wrap(err, "Gagal konek ke LDAP.")
	}
	defer conn.Close()

	if cfg.LDAP_USE_TLS {
		err = conn.StartTLS(&tls.Config{InsecureSkipVerify: true})
		if err != nil {
			return "", errors.Wrap(err, "TLS gagal diaktifkan.")
		}
	}

	err = conn.Bind(dn, request.Password)
	if err != nil {
		return "", errors.Wrap(err, "Username atau Password salah.")
	}

	// ✅ if LDAP success → generate token
	user, err := s.AuthRepository.FindUserByUsername(ctx, request.Username)
	if err != nil {
		return "", errors.Wrap(err, "invalid username")
	}

	token, err := helper.GenerateToken(ctx, user, "access_token", time.Now())
	if err != nil {
		return "", errors.Wrap(err, "gagal generate token")
	}

	return token, nil
}

func (s *AuthServiceImpl) CheckLDAP(ctx context.Context, request dto.LoginRequest) error {
	cfg, err := helper.GetConfig()
	if err != nil {
		return errors.Wrap(err, "gagal load config")
	}

	ldapURL := fmt.Sprintf("%s:%d", cfg.LDAP_SERVER, cfg.LDAP_PORT)
	dn := fmt.Sprintf("uid=%s,%s,%s", ldap.EscapeFilter(request.Username), cfg.LDAP_USER_DN, cfg.LDAP_BASE_DN)

	conn, err := ldap.DialURL(ldapURL)
	if err != nil {
		return errors.Wrap(err, "Gagal konek ke LDAP.")
	}
	defer conn.Close()

	if cfg.LDAP_USE_TLS {
		err = conn.StartTLS(&tls.Config{InsecureSkipVerify: true})
		if err != nil {
			return errors.Wrap(err, "TLS gagal diaktifkan.")
		}
	}

	err = conn.Bind(dn, request.Password)
	if err != nil {
		return errors.Wrap(err, "Password Invalid.")
	}

	return nil
}
