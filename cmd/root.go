package cmd

import (
	"fmt"
	"github.com/dennisstritzke/trex_exporter/exporter"
	"github.com/spf13/cobra"
	"os"
)

const (
	flagTrexApiAddress   = "api-address"
	flagWorker           = "worker"
	flagWebListenAddress = "web.listen-address"
)

var Version string
var RootCmd = &cobra.Command{
	Use:     "trex_exporter",
	Short:   "Prometheus exporter for the T-Rex NVIDIA GPU miner.",
	Long:    "",
	Run:     defaultCommand,
	Version: Version,
}

func init() {
	envAddr := os.Getenv("TREX_EXPORTER_PORT")
	if envAddr == "" {
		envAddr = "9788"
	}
	envAddr = "0.0.0.0:" + envAddr
	envMinerurl := os.Getenv("TREX_MINER_URL")
	if envMinerurl == "" {
		envMinerurl = "http://localhost:4067"
	}
	envWorkername := os.Getenv("TREX_WORKER_NAME")
	if envWorkername == "" {
		envWorkername = "trex"
	}
	RootCmd.PersistentFlags().StringVar(&exporter.TrexApiAddress, flagTrexApiAddress, envMinerurl,
		"Address of the t-rex API.")

	RootCmd.PersistentFlags().StringVar(&exporter.Worker, flagWorker, envWorkername,
		"Name to identify the T-Rex Miner. The name will be included in every metric as the label 'worker'.")

	RootCmd.PersistentFlags().StringVar(&exporter.WebListenAddress, flagWebListenAddress, envAddr,
		"Address on which to expose metrics.")
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func defaultCommand(_ *cobra.Command, _ []string) {
	exporter.Serve()
}
