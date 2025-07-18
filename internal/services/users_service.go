package services

import (
	"context"
	"errors"

	"github.com/ecbDeveloper/go-bid/internal/store/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

var ErrDuplciateEmailOrPassword = errors.New("username or email already exists")

type UserService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

func NewUserService(pool *pgxpool.Pool) UserService {
	return UserService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

func (us *UserService) CreateUser(ctx context.Context, userName, email, password, bio string) (uuid.UUID, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return uuid.UUID{}, err
	}

	args := pgstore.CreateUserParams{
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
			return uuid.UUID{}, ErrDuplciateEmailOrPassword
		}

		return uuid.UUID{}, err
	}

	return id, nil
}
