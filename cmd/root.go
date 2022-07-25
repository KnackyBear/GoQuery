package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

type config struct {
	Url       []string // Urls to request
	Csv       string   // CSV filename
	HasHeader bool     // CSV has header ?
	Method    string   // HTTP Method
	Thread    int      // Number of thread allow
	Delay     int      // Delay time in ms between each request
}

var (
	currentConfig config // Current configuration
)

var RootCmd = &cobra.Command{
	Use:                   "goquery [get|post|put|delete|head] --url=[URL] --csv=[PATH/TO/CSV] (--thread=[THREADS]) (--delay=[DELAY IN MS])",
	Short:                 "Multi-Query manager",
	Long:                  `This application manage multi-query with variables in CSV.`,
	ValidArgs:             []string{"url", "csv", "thread", "delay"},
	DisableFlagsInUseLine: true,
	CompletionOptions: cobra.CompletionOptions{
		DisableDescriptions: true,
	},
}

func init() {
	cobra.OnInitialize()
	RootCmd.PersistentFlags().StringArrayVar(&currentConfig.Url, "url", nil, "Url to request")
	RootCmd.PersistentFlags().StringVar(&currentConfig.Csv, "csv", "", "Path to CSV")
	RootCmd.PersistentFlags().IntVar(&currentConfig.Thread, "thread", 1, "Number of threads")
	RootCmd.PersistentFlags().IntVar(&currentConfig.Delay, "delay", 0, "Insert delay between each request for the same url")
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
