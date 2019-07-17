package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"gotest/util"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "just run",
	Long:  `just run`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Println(util.WrapRed("请输入要执行的Test方法!"))
			cmd.Help()
			return
		}

		// 拼接go test 命令
		funcName := strings.TrimSpace(args[0])
		testFunc := caseSuite.GetFunc(funcName)
		verbose, _ := cmd.Flags().GetBool("verbose")
		count, _ := cmd.Flags().GetInt("count")
		i := strings.LastIndex(testFunc.Path, "/")
		cmdStr := fmt.Sprintf(`cd %s && go test -run %s -count=%d -v=%v .`, testFunc.Path[:i], funcName, count, verbose)

		// 执行go test
		stdout, stderr := util.CMD(cmdStr)
		cmd.Println(stdout + stderr)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().IntP("count", "c", 1, "执行次数")
	runCmd.Flags().BoolP("verbose", "v", true, "是否打印详细信息")

}
