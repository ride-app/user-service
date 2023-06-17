package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/ride-app/entity-service/api/gen/ride/entity/v1alpha1/entityv1alpha1connect"
	"github.com/ride-app/entity-service/config"
	"github.com/ride-app/entity-service/di"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetReportCaller(true)
	if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) != 0 {
		log.SetFormatter(&log.TextFormatter{
			DisableLevelTruncation: true,
			PadLevelText:           true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				dir, err := os.Getwd()
				if err != nil {
					dir = ""
				} else {
					dir = dir + "/"
				}

				filename := strings.Replace(f.File, dir, "", -1)

				return fmt.Sprintf("(%s)", path.Base(f.Function)), fmt.Sprintf(" %s:%d", filename, f.Line)

			},
		})
	}

	err := cleanenv.ReadEnv(&config.Env)

	if err != nil {
		log.Warnf("Could not load config: %v", err)
	}
}

func main() {
	service, err := di.InitializeService()

	if err != nil {
		log.Fatalf("Failed to initialize service: %v", err)
	}

	log.Info("Service Initialized")

	path, handler := entityv1alpha1connect.NewEntityServiceHandler(service)
	mux := http.NewServeMux()
	mux.Handle(path, handler)

	panic(http.ListenAndServe(
		fmt.Sprintf("0.0.0.0:%d", config.Env.Port),
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	))

}
