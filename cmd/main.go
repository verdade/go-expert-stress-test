package main

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/verdade/go-expert-stress-test/internal/httpstress"
)

var (
	urlFlag         string
	requestsFlag    int
	concurrencyFlag int
)
var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "stress test construction tool",
	Long:  "stress test construction tool",
	Run: func(cmd *cobra.Command, args []string) {

		if _, err := url.ParseRequestURI(urlFlag); err != nil {
			panic(err)
		}

		if requestsFlag == 0 || concurrencyFlag == 0 {
			panic("Flags with zero values ")
		}

		hs := httpstress.NewHttpStresser()
		err := hs.StartTest(urlFlag, int(requestsFlag), int(concurrencyFlag))
		if err != nil {
			fmt.Println(err)
		}
	},
}

func main() {

	rootCmd.Flags().StringVarP(&urlFlag, "url", "u", "", "URL para fazer requisições")
	rootCmd.Flags().IntVarP(&requestsFlag, "requests", "r", 0, "Número total de requisições a serem feitas")
	rootCmd.Flags().IntVarP(&concurrencyFlag, "concurrency", "c", 0, "Número de requisições concorrentes a serem feitas")

	cobra.CheckErr(rootCmd.Execute())
}
