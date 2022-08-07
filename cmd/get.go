package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/knackybear/goquery/pkg/utils"
	"github.com/spf13/cobra"
)

var cmdGet = &cobra.Command{
	Use:                   "get",
	Short:                 "HTTP Get",
	DisableFlagsInUseLine: true,
	ValidArgsFunction:     cobra.NoFileCompletions,
	RunE: func(cmd *cobra.Command, args []string) error {
		rc := make(chan *http.Response)
		queue := []string{}
		
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
				queue = append(queue, url)
			}
		}

		var wg sync.WaitGroup
		wg.Add(len(queue))

		for idx, req := range queue {
			go func()
		}
		return nil
	},
}

func GetRequest(url string, rc chan *http.Response) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		rc <- resp
	}
}

func init() {
	RootCmd.AddCommand(cmdGet)
}
