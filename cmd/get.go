package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/knackybear/goquery/pkg/utils"
	"github.com/spf13/cobra"
)

var cmdGet = &cobra.Command{
	Use:                   "get",
	Short:                 "HTTP Get",
	DisableFlagsInUseLine: true,
	ValidArgsFunction:     cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		file, err := os.Open(currentConfig.Csv)
		if err != nil {
			return err
		}
		defer file.Close()
		vars := utils.CSVToMap(file)
		for _, variable := range vars {
			for _, url := range currentConfig.Url {
				for k, v := range variable {
					url = strings.ReplaceAll(url, fmt.Sprintf("{%s}", k), v)
				}
				println("GET ", url)
			}
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(cmdGet)
}
