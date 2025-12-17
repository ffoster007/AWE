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

// StartCmd คือ command สำหรับ start server
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
	Long:  `Start the server with specified configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		host, _ := cmd.Flags().GetString("host")
		verbose, _ := cmd.Flags().GetBool("verbose")

		if verbose {
			fmt.Println("Starting server in verbose mode...")
		}

		fmt.Printf("🚀 Server starting at %s:%d\n", host, port)
		// เพิ่มโค้ดสำหรับ start server ของคุณที่นี่
	},
}

// EncryptCmd สำหรับ encryption (ตัวอย่าง)
var EncryptCmd = &cobra.Command{
	Use:   "encrypt [file]",
	Short: "Encrypt a file",
	Long:  `Encrypt a file using AWE encryption algorithm`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		output, _ := cmd.Flags().GetString("output")

		fmt.Printf("🔐 Encrypting file: %s\n", filename)
		if output != "" {
			fmt.Printf("📝 Output: %s\n", output)
		}
		// เพิ่มโค้ด encryption logic ที่นี่
	},
}

// DecryptCmd สำหรับ decryption (ตัวอย่าง)
var DecryptCmd = &cobra.Command{
	Use:   "decrypt [file]",
	Short: "Decrypt a file",
	Long:  `Decrypt a file using AWE decryption algorithm`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		output, _ := cmd.Flags().GetString("output")

		fmt.Printf("🔓 Decrypting file: %s\n", filename)
		if output != "" {
			fmt.Printf("📝 Output: %s\n", output)
		}
		// เพิ่มโค้ด decryption logic ที่นี่
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

	// Flags สำหรับ start command
	StartCmd.Flags().IntP("port", "p", 8080, "Port to start server on")
	StartCmd.Flags().StringP("host", "H", "localhost", "Host address")
	StartCmd.Flags().BoolP("verbose", "v", false, "Enable verbose output")

	// Flags สำหรับ encrypt/decrypt commands
	EncryptCmd.Flags().StringP("output", "o", "", "Output file path")
	DecryptCmd.Flags().StringP("output", "o", "", "Output file path")

	// เพิ่ม subcommands เข้าไปใน root command
	RootCmd.AddCommand(StartCmd)
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
