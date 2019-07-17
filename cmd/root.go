package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"gotest/model"
	"gotest/util"
)

var (
	caseSuite *model.CaseSuite
)

func init() {
	rootDir := util.CurrentDir()
	caseSuite = model.NewCaseSuite(rootDir)
}

var rootCmd = &cobra.Command{
	Use:   "gotest",
	Short: "go test 帮助工具",
	Long:  `go test 帮助工具`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
