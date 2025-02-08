package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var postCmd = &cobra.Command{
	Use:   "post [URL] -d '[JSON]'",
	Short: "Send a POST request with JSON data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		data, _ := cmd.Flags().GetString("data")
		resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(data)))
		if err != nil {
			fmt.Println("Error : ", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	},
}

func init() {
	postCmd.Flags().StringP("data", "d", "{}", "JSON data to send")
	rootCmd.AddCommand(postCmd)
}
