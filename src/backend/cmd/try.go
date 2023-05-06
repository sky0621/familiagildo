/*
Copyright © 2023 sky0621 <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/sky0621/familiagildo/domain/service"
	"github.com/spf13/cobra"
)

// tryCmd represents the try command
var tryCmd = &cobra.Command{
	Use:   "try",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("try called")
		fmt.Println(service.CreateToken())
		fmt.Println(service.CreateAcceptedNumber())
		edt := service.CreateGuestTokenExpirationDate().ToVal()
		fmt.Println(edt)
		fmt.Println(service.ToMailFormattedDatetime(edt))
	},
}

func init() {
	rootCmd.AddCommand(tryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
