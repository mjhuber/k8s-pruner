package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/klog"
)

var (
	// VERSION is set during build
	VERSION string
)

func init() {
	klog.InitFlags(nil)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	flag.Set("alsologtostderr", "true")
	flag.Set("logtostderr", "true")
	flag.Parse()
}

var rootCmd = &cobra.Command{
	Use:   "gns",
	Short: "gns",
	Long:  "A tool to easily switch google cloud projects and gke clusters.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Error, you must specify a sub-command.")
		err := cmd.Help()
		if err != nil {
			fmt.Printf("\n%v\n", err)
		}
		os.Exit(1)
	},
}

// Execute is the main entry point into the command
func Execute(version string) {
	VERSION = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
