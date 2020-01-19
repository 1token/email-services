package cmd

import (
	"context"
	"fmt"
	"github.com/1token/email-services/pkg/config"
	"github.com/1token/email-services/repository"
	"github.com/1token/email-services/session"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net/http"
	"os"
	"strings"

	pb "github.com/1token/email-services/email-apis/generated/go"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:     "serve --config [ config file ]",
	Short:   "Start serving requests.",
	Long:    ``,
	Example: "email-services serve --config examples/config-dev.yaml",
	Run: func(cmd *cobra.Command, args []string) {
		if err := serve(); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve() error {
	// unmarshal config into Struct
	var c config.Config
	err := viper.Unmarshal(&c)
	if err != nil {
		return fmt.Errorf("unmarshal failed: %v", err)
	}

	if err := c.Validate(); err != nil {
		return err
	}

	var (
		grpcServer    *grpc.Server
		wrappedServer *grpcweb.WrappedGrpcServer
	)

	if c.Web.TLSCert != "" {
		creds, err := credentials.NewServerTLSFromFile(c.Web.TLSCert, c.Web.TLSKey)
		if err != nil {
			log.Fatalf("Failed while obtaining TLS certificates. Error: %+v", err)
		}

		/*opts := []grpcrecovery.Option{
			grpcrecovery.WithRecoveryHandler(GrpcRecoveryHandlerFunc),
		}*/

		grpcServer = grpc.NewServer(
			grpc.Creds(creds),
			grpc.StreamInterceptor(grpcmiddleware.ChainStreamServer(
				StreamAuthInterceptor,
				// grpcrecovery.StreamServerInterceptor(opts...),
				grpcrecovery.StreamServerInterceptor(),
			)),
			grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
				UnaryAuthInterceptor,
				// grpcrecovery.UnaryServerInterceptor(opts...),
				grpcrecovery.UnaryServerInterceptor(),
			)),
		)
	}

	pb.RegisterDraftServiceServer(grpcServer, &pb.DraftServiceServer{})
	// pb.RegisterAuthServer(grpcServer, &oidc.UserInfoImpl{})
	// pb.RegisterUsersServer(grpcServer, &impl.UserServerImpl{db})
	// pb.RegisterJmsApiServer(grpcServer, &impl.JmapServerImpl{db})
	// pb.RegisterGmailServer(grpcServer, &impl.GmailServerImpl{db})
	// pb.RegisterGmailServer(grpcServer, &impl.LabelsServerImpl{db})
	// pb.RegisterFilesServer(grpcServer, &impl.FileServerImpl{db})

	wrappedServer = grpcweb.WrapServer(grpcServer,
		grpcweb.WithOriginFunc(func(origin string) bool {
			for _, allowedOrigin := range c.Web.AllowedOrigins {
				if allowedOrigin == "*" || origin == allowedOrigin {
					return true
				}
			}
			return false
		}),
	)

	restMux := http.NewServeMux() // http.DefaultServeMux

	httpServer := http.Server{
		Addr: c.Web.Addr(),
		Handler: http.HandlerFunc(
			func(resp http.ResponseWriter, req *http.Request) {
				if req.Method == http.MethodOptions {
					for _, allowedOrigin := range c.Web.AllowedOrigins {
						if allowedOrigin == "*" || allowedOrigin == req.Header.Get("Origin") {
							SetCors(resp, allowedOrigin)
							return
						}
					}
					return
				}

				if IsGrpcRequest(req) {
					if wrappedServer.IsGrpcWebRequest(req) {
						wrappedServer.ServeHTTP(resp, req)
					} else {
						grpcServer.ServeHTTP(resp, req)
					}
				} else {
					restMux.ServeHTTP(resp, req)
				}
			},
		),
	}

	// restMux.HandleFunc(common.AUTH_SIGNIN_PATH, oidc.SignIn())
	// restMux.HandleFunc(common.AUTH_SIGNOUT_PATH, oidc.SignOut())
	// restMux.HandleFunc(common.AUTH_REDIRECT_PATH, oidc.HandleCallback())
	restMux.HandleFunc(c.File.UploadApi, session.AuthorizeRest(repository.FileUploadHandler))

	// make a channel for each server to fail properly
	errc := make(chan error, 1)

	go func() {
		errc <- httpServer.ListenAndServeTLS(c.Web.TLSCert, c.Web.TLSKey)
		close(errc)
	}()

	return <-errc
}

func SetCors(w http.ResponseWriter, allowedOrigin string) {
	w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, x-grpc-web, sessionid, x-user-agent, authorization-token")
	w.Header().Set("Access-Control-Max-Age", "600")
}

func IsGrpcRequest(req *http.Request) bool {
	return strings.Contains(req.Header.Get("Content-Type"), "application/grpc")
}

/*func GrpcRecoveryHandlerFunc(p interface{}) error {
	fmt.Printf("p: %+v\n", p)
	return status.Errorf(codes.Internal, "Unexpected error")
}*/

func StreamAuthInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	ctx, err := session.AuthorizeGrpc(stream.Context())
	if err != nil {
		log.Errorf("%s %v", info.FullMethod, err)
		return err
	}

	wrapped := grpcmiddleware.WrapServerStream(stream)
	wrapped.WrappedContext = ctx
	return handler(srv, wrapped)
}

func UnaryAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx, err := session.AuthorizeGrpc(ctx)
	if err != nil {
		log.Errorf("%s %v", info.FullMethod, err)
		return nil, err
	}

	return handler(ctx, req)
}
