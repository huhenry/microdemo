package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/huhenry/microdemo/pkg/server"
	"github.com/spf13/cobra"
	klog "k8s.io/klog/v2"
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

	klog.InitFlags(flag.CommandLine)
	//klog.InitFlags(nil)
	flag.Set("logtostderr", "true")
	// make sure we flush before exiting
	defer klog.Flush()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
