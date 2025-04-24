package cmd

import (
	"github.com/spf13/cobra"
	"kryptonim-interview/app"
	"kryptonim-interview/app/errs"
	"kryptonim-interview/app/routers"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "run api",
	Run:   ApiCommandRun,
}

func ApiCommandRun(cmd *cobra.Command, args []string) {
	container := app.NewContainer()
	router := routers.SetupCore(container)
	err := router.Run("0.0.0.0", "8080")
	if err != nil {
		errs.FatalOnError(err, "[13187187] cannot start core api")
	}
}
