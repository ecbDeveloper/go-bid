package services

import (
	"context"
	"errors"

	"github.com/ecbDeveloper/go-bid/internal/db/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplciateEmailOrUsername = errors.New("username or email already exists")
	ErrInvalidCredentials       = errors.New("invalid credential")
)

type UserService struct {
	pool    *pgxpool.Pool
	queries *sqlc.Queries
}

func NewUserService(pool *pgxpool.Pool) UserService {
	return UserService{
		pool:    pool,
		queries: sqlc.New(pool),
	}
}

func (us *UserService) CreateUser(ctx context.Context, userName, email, password, bio string) (uuid.UUID, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return uuid.UUID{}, err
	}

	args := sqlc.CreateUserParams{
		UserName:     userName,
		Email:        email,
		PasswordHash: hashedPassword,
		Bio:          bio,
	}

	id, err := us.queries.CreateUser(ctx, args)
	if err != nil {
		var pgErr *pgconn.PgError
		// errors.Is verifica se o erro é exatamente o mesmo do target
		// errors.As verifica se o erro é do mesmo tipo que o tipo indicado,
		// e caso seja, conseguimos acessar as propriedades desse tipo de erro
		if errors.As(err, *pgErr) && pgErr.Code == "23505" {
			return uuid.UUID{}, ErrDuplciateEmailOrUsername
		}

		return uuid.UUID{}, err
	}

	return id, nil
}

func (us *UserService) AuthenticateUser(ctx context.Context, email, password string) (uuid.UUID, error) {
	user, err := us.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return uuid.UUID{}, ErrInvalidCredentials
		}
		return uuid.UUID{}, err
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return uuid.UUID{}, ErrInvalidCredentials
		}
		return uuid.UUID{}, err
	}

	return user.ID, nil
}
