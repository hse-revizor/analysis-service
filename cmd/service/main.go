package main

import (
	"fmt"

	"github.com/hse-revizor/analysis-service/internal/di"
	"github.com/hse-revizor/analysis-service/internal/utils/config"
	"github.com/hse-revizor/analysis-service/internal/utils/flags"
	"github.com/slipneff/gogger"
	goggerlog "github.com/slipneff/gogger/log"
)

// @title           Analysis Service API
// @version         1.0
// @description     This is a Swagger documentation.
// @termsOfService  http://swagger.io/terms/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:8080
// @BasePath /api/v1
func main() {
	flags := flags.MustParseFlags()
	cfg := config.MustLoadConfig(flags.EnvMode)
	container := di.New(cfg)
	gogger.ConfigureZeroLogger()

	goggerlog.Info(fmt.Sprintf("Server starting on %s:%d", cfg.Host, cfg.Port))

	err := container.GetHttpServer().ListenAndServe()
	if err != nil {
		goggerlog.Panic(err, "Fail serve HTTP")
	}
}
