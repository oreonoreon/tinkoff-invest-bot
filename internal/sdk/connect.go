package sdk

import (
	"context"
	"crypto/tls"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/metadata"
	"time"
	"tinkoff-invest-bot/internal/config"
)

func Connection() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("invest-public-api.tinkoff.ru:443",
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithPerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: config.Token,
		})))

	return conn, err
}

func NewContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	ctx = metadata.AppendToOutgoingContext(ctx, "x-app-name", config.Appname)
	return ctx, cancel
}
func NewContextStream() context.Context {
	ctx := context.TODO()
	ctx = metadata.AppendToOutgoingContext(ctx, "x-app-name", config.Appname)
	return ctx
}
