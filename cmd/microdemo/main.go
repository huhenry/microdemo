package microdemo

import (
	"fmt"
	"os"

	"github.com/huhenry/microdemo/pkg/server"
	"github.com/spf13/cobra"
)

var (
	port int
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "microdemo",
		Short: "microdemo a super simple service for generating different HTTP codes.",
		Long:  `microdemo a super simple service for generating different HTTP codes.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return server.Start(port)
		},
	}
	rootCmd.Flags().IntVarP(&port, "port", "o", 9696, "listen port")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
