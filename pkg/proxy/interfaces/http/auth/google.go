package auth

import (
	"net/http"

	"github.com/vardius/go-api-boilerplate/pkg/common/application/errors"
	"github.com/vardius/go-api-boilerplate/pkg/common/application/http/response"
	user_proto "github.com/vardius/go-api-boilerplate/pkg/user/infrastructure/proto"
	user_grpc "github.com/vardius/go-api-boilerplate/pkg/user/interfaces/grpc"
)

type google struct {
	userClient user_proto.UserClient
	authClient user_proto.AuthClient
}

func (g *google) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	accessToken := r.FormValue("accessToken")
	data, e := getProfile(accessToken, "https://www.googleapis.com/oauth2/v2/userinfo")
	if e != nil {
		response.WithError(r.Context(), errors.Wrap(e, "Invalid access token", errors.INVALID))
		return
	}

	_, e = g.userClient.DispatchCommand(r.Context(), &user_proto.DispatchCommandRequest{
		Name:    user_grpc.RegisterUserWithGoogle,
		Payload: data,
	})

	if e != nil {
		response.WithError(r.Context(), errors.Wrap(e, "Invalid request", errors.INVALID))
		return
	}

	token, e := f.authClient.DispatchCommand(r.Context(), &auth_proto.GetToken{
		Email: data.Email,
	})

	if e != nil {
		response.WithError(r.Context(), errors.Wrap(e, "Generate token failure", errors.INTERNAL))
		return
	}

	response.WithPayload(r.Context(), &authTokenResponse{token})
	return
}

// NewGoogle creates google auth handler
func NewGoogle(u user_proto.UserClient, a auth_proto.AuthClient) http.Handler {
	return &google{u, a}
}