package cmd

import (
	"fmt"
	"go-courseware/api"

	"github.com/spf13/cobra"
)

var (
	containerID string
	publishID   string
	jwtToken    string
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch courseware data from the API",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := api.FetchCourseware(containerID, publishID, jwtToken)
		if err != nil {
			fmt.Println("Error fetching courseware:", err)
			return
		}
		fmt.Println("Courseware Data:")
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().StringVarP(&containerID, "container", "c", "", "Container ID (required)")
	fetchCmd.Flags().StringVarP(&publishID, "publish", "p", "", "Publish ID (required)")
	fetchCmd.Flags().StringVarP(&jwtToken, "jwt", "j", "", "JWT Token (required)")

	fetchCmd.MarkFlagRequired("container")
	fetchCmd.MarkFlagRequired("publish")
	fetchCmd.MarkFlagRequired("jwt")
}
