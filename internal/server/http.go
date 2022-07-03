package server

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/gorilla/handlers"
	v1 "kratos-realworld/api/realworld/v1"
	"kratos-realworld/internal/conf"
	"kratos-realworld/internal/errorhandler"
	"kratos-realworld/internal/pkg/middleware/auth"
	"kratos-realworld/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	nethttp "net/http"
)

// NewSkipRoutersMatcher 通过正则匹配取消jwt token验证
func NewSkipRoutersMatcher() selector.MatchFunc {

	allowList := make(map[string]struct{})
	// /包名.服务名/方法名
	allowList["/realworld.v1.Realworld/Login"] = struct{}{}
	allowList["/realworld.v1.Realworld/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := allowList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, jwt *conf.JWT, realWorlder *service.RealWorldService, logger log.Logger) *http.Server {
	fmt.Println("logger: ", logger)
	var opts = []http.ServerOption{
		http.ErrorEncoder(errorhandler.ErrorEncoder),
		http.Middleware(
			recovery.Recovery(),
			//tracing.Server(),
			//logging.Server(logger),
			//auth.JWTAuth(jwt.Token),

			selector.Server(auth.JWTAuth(jwt.Token)).Match(NewSkipRoutersMatcher()).Build(),
		),
		http.Filter(
			func(h nethttp.Handler) nethttp.Handler {
				return nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
					fmt.Println("route filter in")
					h.ServeHTTP(w, r)
					fmt.Println("route filter out")
				})
			},
			handlers.CORS(
				handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
				handlers.AllowedOrigins([]string{"*"}),
			),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterRealworldHTTPServer(srv, realWorlder)
	return srv
}
