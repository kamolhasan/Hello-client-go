package cmd

import (
	"github.com/kamolhasan/Hello-client-go/book"
	"github.com/spf13/cobra"
	"log"
)

var deleteCmd = &cobra.Command{
	Use:"delete",

}

var service3Cmd = &cobra.Command{
	Use:"service",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Deleting Service")
		book.DeleteService()
	},
}


var deployment3Cmd = &cobra.Command{
	Use:"deployment",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Deleting demployment")
		book.DeleteDeployment()
	},
}


var ingress3Cmd = &cobra.Command{
	Use:"ingress",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Deleting Ingress")
		book.DeleteIngress()
	},
}


func init(){
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.AddCommand(service3Cmd)
	deleteCmd.AddCommand(deployment3Cmd)
	deleteCmd.AddCommand(ingress3Cmd)
}