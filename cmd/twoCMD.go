package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdTwoCmd = &cobra.Command{
	Use: "CmdONE",
	Short: "bu qisa tesvirdir komandanin",
	Long: "Bu uzun tesvuridir comandaninfsdfsdfs df sd fsd f sdfsdf",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cmd Two cagrildi")
		number, _ := cmd.Flags().GetInt("number")
		fmt.Println("istifade edilecek nomre: ", number)

		fmt.Println("istifade edilecek nomrenin kvadrati: ", number*number)
	},

}

func init(){
	rootCmd.AddCommand(cmdOneCmd)
	cmdOneCmd.Flags().Int("number", 0, "A help for number")
}
