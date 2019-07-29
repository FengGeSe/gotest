package cmd

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/FengGeSe/gotest/util"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "列出本项目所有case",
	Long:  `列出本项目所有case`,
	Run: func(cmd *cobra.Command, args []string) {
		currentDir := util.CurrentDir()

		verbose, _ := cmd.Flags().GetBool("verbose")
		if verbose {
			for path, caseSlc := range caseSuite.GetCaseArray() {
				pathStr := "." + strings.TrimPrefix(path, currentDir)
				cmd.Println(util.WrapBlue(pathStr))
				for _, f := range caseSlc {
					cmd.Println("\t" + f.String())
				}
			}
		} else {
			// 简单打印
			for _, caseSlc := range caseSuite.GetCaseArray() {
				for _, f := range caseSlc {
					cmd.Println(f.String())
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("verbose", "v", false, "详细打印出Test函数所在的文件")
}
