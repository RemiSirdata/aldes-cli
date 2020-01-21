package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(routerCmd)
}

var routerCmd = &cobra.Command{
	Use:   "products",
	Short: "List Aldes products",
	Long:  `PRODUCTS list all Aldes products`,

	Run: listProducts,
}

func listProducts(cmd *cobra.Command, args []string) {
	client, err := getAldesClient()
	if err != nil {
		log.Fatalf("fail to init: %s", err.Error())
	}
	products, err := client.GetProducts()
	if err != nil {
		log.Fatalf("fail to fetch products: %s", err.Error())
	}
	log.Printf("products: %+v", products)
}
