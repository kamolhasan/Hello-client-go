package cmd

import (
	"github.com/kamolhasan/Hello-client-go/book"
	"log"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use: "create",
}

var serviceCmd = &cobra.Command{
	Use: "service",

	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Creating Service")
		book.CreateService()
	},
}

var deploymentCmd = &cobra.Command{
	Use: "deployment",

	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Creating Deployment")
		book.CreateDeployment()

	},
}

var ingressCmd = &cobra.Command{
	Use: "ingress",

	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Creating ingress service")
		book.CreateIngress()

	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(serviceCmd)
	createCmd.AddCommand(deploymentCmd)
	createCmd.AddCommand(ingressCmd)
}
