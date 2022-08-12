package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)
var cmdOneCmd = &cobra.Command{
	Use: "CmdTWO",
	Short: "bu qisa tesvirdir komandanin",
	Long: "Bu uzun tesvuridir comandaninfsdfsdfs df sd fsd f sdfsdf",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cmd One Called")
		number, _ := cmd.Flags().GetInt("number")
		fmt.Println("istifade edilecek nomre: ", number)

		fmt.Println("istifade edilecek nomrenin kvadrati: ", number*number)
	},

}

func init(){
	rootCmd.AddCommand(cmdTwoCmd)
	cmdTwoCmd.Flags().Int("number2", 0, "A help for number")
}



