package cmd

import (
	"fmt"
	"github.com/RemiSirdata/aldes-cli/api"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "aldes",
	Short: "Aldes CLI allows you to interact with Aldes products APIs",
	Long: `Aldes CLI allows you to interact with Aldes product API
                Complete documentation is available at https://github.com/RemiSirdata/aldes-cli`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getAldesClient() (*api.Aldes, error) {
	return api.New(os.Getenv("ALDES_LOGIN"), os.Getenv("ALDES_PASSWORD"))
}
