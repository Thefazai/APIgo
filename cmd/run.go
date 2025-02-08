package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type Request struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers,omitempty"`
	Body    string            `json:"body,omitempty"`
}

var runCmd = &cobra.Command{
	Use:   "run [file]",
	Short: "Run an API test script",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := os.ReadFile(args[0])
		var requests []Request
		json.Unmarshal(file, &requests)

		for _, req := range requests {
			fmt.Printf("Executing %s %s...\n", req.Method, req.URL)
			httpReq, _ := http.NewRequest(req.Method, req.URL, nil)
			client := &http.Client{}
			resp, _ := client.Do(httpReq)
			body, _ := io.ReadAll(resp.Body)
			fmt.Println(string(body))
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
