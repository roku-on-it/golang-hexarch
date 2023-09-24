package driven

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"hexarch/core/domain"
	"hexarch/core/ports"
	"strconv"
	"strings"
)

type userRedisAdapter struct {
	c *redis.Client
}

func NewUserRedisAdapter(addr string) *userRedisAdapter {
	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &userRedisAdapter{c}
}

func (a *userRedisAdapter) FindUserByID(id string) (*domain.User, error) {
	// TODO: Implement
	return nil, nil
}

func (a *userRedisAdapter) FindUserByUsername(u string) (*domain.User, error) {
	ctx := context.Background()

	keys, _, _ := a.c.Scan(ctx, 0, "user:"+u+":*", 0).Result()

	if len(keys) == 0 {
		return nil, ports.ErrRecordNotFound
	}

	res, _ := a.c.Get(ctx, keys[0]).Result()
	splt := strings.Split(res, ";")

	age, _ := strconv.Atoi(splt[3])

	return &domain.User{splt[0], splt[2], splt[2], age}, nil
}

func (a *userRedisAdapter) CreateUser(input domain.CreateUserInput) (*domain.User, error) {
	_, err := a.FindUserByUsername(input.Username)

	if err == nil {
		if !errors.Is(err, ports.ErrRecordNotFound) {
			return nil, ports.ErrDuplicatedKey
		}

		return nil, err
	}

	ctx := context.Background()

	id := uuid.New().String()

	_, err = a.c.Set(ctx, "user:"+input.Username+":"+id, id+";"+input.Username+";"+input.FullName+";"+strconv.Itoa(input.Age), 0).Result()

	if err != nil {
		return nil, err
	}

	return &domain.User{id, input.Username, input.FullName, input.Age}, nil
}
