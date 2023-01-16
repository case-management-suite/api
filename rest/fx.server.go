package rest

import (
	// "dig"
	"context"
	"flag"
	"log"

	"github.com/case-management-suite/api/apicommon"
	"github.com/case-management-suite/api/rest/restapi"
	"github.com/case-management-suite/api/rest/restapi/ops"
	"github.com/case-management-suite/common/config"
	"go.uber.org/fx"
)

func NewRestServer(port *int) func(lc fx.Lifecycle, api ops.CaseMgmtAPI) *restapi.Server {
	return func(lc fx.Lifecycle, api ops.CaseMgmtAPI) *restapi.Server {
		srv := restapi.NewServer(&api)
		srv.Port = *port

		// srv.TLSCertificate = flags.Filename("./sample_ssl/certificate.crt")
		// srv.TLSCertificateKey = flags.Filename("./sample_ssl/key.key")

		lc.Append(fx.Hook{
			OnStart: func(_ context.Context) error {
				// serve API
				go func() {
					if err := srv.Serve(); err != nil {
						log.Fatalln(err)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return srv.Shutdown()
			},
		})

		return srv
	}
}

func CreateLiteAPIServer(appConfig config.AppConfig) *fx.App {
	var portFlag = flag.Int("port", 8080, "Port to run this service on")
	return fx.New(
		apicommon.FxGetClientServices(appConfig.RulesServiceConfig),
		fx.Provide(
			NewCaseMgmtAPI,
			NewRestServer(portFlag),
		), fx.Invoke(func(*restapi.Server) {}))
}

func CreateLiteTestAPIServer(appConfig config.AppConfig) *fx.App {
	if flag.Lookup("port") == nil {
		flag.Int("port", 8080, "Port to run this service on")
	}
	portFlag := flag.Lookup("port").Value.(flag.Getter).Get().(int)
	return fx.New(
		config.FxConfig(appConfig),
		apicommon.FxGetClientServices(appConfig.RulesServiceConfig),
		fx.Provide(
			NewCaseMgmtAPI,
			NewRestServer(&portFlag),
		// NewHTTPServer(":8080"),
		), fx.Invoke(func(_ *restapi.Server, lc fx.Lifecycle) {

		}))
}
