package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [URL]",
	Short: "Send a GET request",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		resp, err := http.Get(url)
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
	rootCmd.AddCommand(getCmd)
}
