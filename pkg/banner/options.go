package banner

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	Version = "1.0.0"
)

// RootCmd คือ command หลัก
var RootCmd = &cobra.Command{
	Use:     "awe [flags]",
	Short:   "AWE - A encryption tool for security testing",
	Long:    `AWE is a encryption tool used for security testing developed by AVACX.`,
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
		// ถ้าไม่มี flag ใดๆ ให้แสดง help
		cmd.Help()
	},
}

func Initbanner() {
	// นำเข้าจาก initialize.go
}

// Execute เรียกใช้ root command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
