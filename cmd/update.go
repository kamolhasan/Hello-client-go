package cmd

import (
	"github.com/kamolhasan/Hello-client-go/book"
	"github.com/spf13/cobra"
	"log"
)

var updateCmd = &cobra.Command{
	Use:"update",

}

var deployment2Cmd= &cobra.Command{
		Use:"deployment",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("updating deployment")
			book.UpdateDeployment()

		},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.AddCommand(deployment2Cmd)
}

