package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "1.0.0"
)

// Root command
var rootCmd = &cobra.Command{
	Use:     "awe [flags] <value>", // แก้ตรงนี้
	Short:   "AWE - Your awesome CLI tool",
	Long:    `AWE is a encryption tool used for security testing developed by AVACX.`,
	Version: version,
}

// Start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
	Long:  `Start the server with specified configuration.`,
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

func init() {
	// ปิดการใช้งานคำสั่ง completion เริ่มต้น
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// กำหนดข้อความ Usage Template แบบกำหนดเอง (ถ้าต้องการ)
	rootCmd.SetUsageTemplate(`Usage:
  {{.UseLine}}{{if .HasAvailableSubCommands}}

Available commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [-h] --help" to see help for each command{{end}}
`)

	// เพิ่ม subcommands เข้าไปใน root command
	rootCmd.AddCommand(startCmd)

	// Global flags
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "Enable debug mode")
	rootCmd.PersistentFlags().BoolP("connect", "c", false, "Connect to AWE Network")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
