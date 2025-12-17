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
	Use:     "awe [flags] <value>",
	Short:   "AWE - Your awesome CLI tool",
	Long:    `AWE is a encryption tool used for security testing developed by AVACX.`,
	Version: Version,
}

// DecryptCmd สำหรับ decryption
var DecryptCmd = &cobra.Command{
	Use:   "decrypt [file]",
	Short: "Decrypt a file",
	Long:  `Decrypt a file using AWE decryption algorithm`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		output, _ := cmd.Flags().GetString("output")

		fmt.Printf("Decrypting file: %s\n", filename)
		if output != "" {
			fmt.Printf("Output: %s\n", output)
		}
		// เพิ่มโค้ด decryption logic ที่นี่
	},
}

// EncryptCmd สำหรับ encryption
var EncryptCmd = &cobra.Command{
	Use:   "encrypt [file]",
	Short: "Encrypt a file",
	Long:  `Encrypt a file using AWE encryption algorithm`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		output, _ := cmd.Flags().GetString("output")

		fmt.Printf("Encrypting file: %s\n", filename)
		if output != "" {
			fmt.Printf("Output: %s\n", output)
		}
		// เพิ่มโค้ด encryption logic ที่นี่
	},
}

func init() {
	// ปิดการใช้งานคำสั่ง completion เริ่มต้น
	RootCmd.CompletionOptions.DisableDefaultCmd = true

	// กำหนดข้อความ Usage Template
	RootCmd.SetUsageTemplate(`Usage:
  {{.UseLine}}{{if .HasAvailableSubCommands}}

Available commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command{{end}}
`)

	// Flags สำหรับ encrypt/decrypt commands
	EncryptCmd.Flags().StringP("output", "o", "", "Output file path")
	DecryptCmd.Flags().StringP("output", "o", "", "Output file path")

	// เพิ่ม subcommands เข้าไปใน root command
	RootCmd.AddCommand(EncryptCmd)
	RootCmd.AddCommand(DecryptCmd)

	// Global flags
	RootCmd.PersistentFlags().BoolP("debug", "d", false, "Enable debug mode")
	RootCmd.PersistentFlags().BoolP("connect", "c", false, "Connect to AWE Network")
}

// Execute เรียกใช้ root command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
