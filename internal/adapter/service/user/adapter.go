package user

import (
	"context"

	"github.com/Ilhamkawe/crowdfunding-proto/protogen/go/auth"
	protoAuth "github.com/Ilhamkawe/crowdfunding-proto/protogen/go/auth"
	"github.com/Ilhamkawe/grpc-campaign-service/internal/application/domain/user"
	"github.com/Ilhamkawe/grpc-campaign-service/internal/port"
	"google.golang.org/grpc"
)

type userServiceAdapter struct {
	client protoAuth.AuthServiceClient
}

func NewUserServiceAdapter(conn *grpc.ClientConn) port.UserServicePort {
	return &userServiceAdapter{
		client: auth.NewAuthServiceClient(conn),
	}
}

func (a *userServiceAdapter) GetUserByID(ctx context.Context, userID int) (serviceport.UserDTO, error) {
	resp, err := a.client.FetchUser(ctx, &auth.SendID{Id: int32(userID)})
	if err != nil {
		return user.User{}, err
	}

	return user.User{
		// ID:     int(resp.Id),
		// Name:   resp.Name,
		// Email:  resp.Email,
		// Avatar: resp.AvatarFileName,
	}, nil
}
