package port

import (
	"context"

	"github.com/Ilhamkawe/grpc-campaign-service/internal/application/domain/user"
)

type UserServicePort interface {
	GetUserByID(ctx context.Context, userID int) (user.User, error)
}
