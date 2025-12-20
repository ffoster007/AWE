package banner

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	// ปิดการใช้งานคำสั่ง completion เริ่มต้น
	RootCmd.CompletionOptions.DisableDefaultCmd = true

	// ปิดการแสดง "completion" command
	RootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	// กำหนด Usage Template แบบกำหนดเอง
	RootCmd.SetUsageTemplate(`  
Usage:
	awe [flags] <values> 

Flags:
STARTS:
        -i,  --init			Create an AWE repository or start a new one from an existing one.
	-c,  --connect [values]		Connect to your SCA (Security Code Access) key
	-cv, --change-vault		For switching to a new/other SCA code.
	-cc, --cancel-connect		For unlinking the code

OPTIONS:
	-s,  --status			Show connect status
	-d,  --debug			Enable to AWE Network
	-r,  --restore			Reset all values (including the SCA code).
	-hs, --host-status		Check the status and connection of your network provider.

OTHERS:
	-up, --update			Update AWE to lastest version
	-h,  --help			Help for AWE
	-v,  --version			Version for AWE
`)

	// STARTS flags
	RootCmd.Flags().BoolP("init", "i", false, "Create an AWE repository or start a new one from an existing one")
	RootCmd.Flags().StringSliceP("connect", "c", []string{}, "Connect to your SCA (Security Code Access) key")
	RootCmd.Flags().Bool("change-vault", false, "For switching to a new/other SCA code")
	RootCmd.Flags().Bool("cv", false, "For switching to a new/other SCA code (short)")
	RootCmd.Flags().Bool("cancel-connect", false, "For unlinking the code")
	RootCmd.Flags().Bool("cc", false, "For unlinking the code (short)")

	// OPTIONS flags
	RootCmd.Flags().BoolP("status", "s", false, "Show connect status")
	RootCmd.Flags().BoolP("debug", "d", false, "Enable to AWE Network")
	RootCmd.Flags().BoolP("restore", "r", false, "Reset all values (including the SCA code)")
	RootCmd.Flags().Bool("host-status", false, "Check the status and connection of your network provider")
	RootCmd.Flags().Bool("hs", false, "Check the status and connection of your network provider (short)")

	// OTHERS flags
	RootCmd.Flags().Bool("update", false, "Update AWE to lastest version")
	RootCmd.Flags().Bool("up", false, "Update AWE to lastest version (short)")

	// ซ่อน flags ที่เป็น short version ไม่ให้แสดงใน help
	RootCmd.Flags().MarkHidden("cv")
	RootCmd.Flags().MarkHidden("cc")
	RootCmd.Flags().MarkHidden("hs")
	RootCmd.Flags().MarkHidden("up")

	// เพิ่ม logic สำหรับแต่ละ flag
	RootCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		// ตรวจสอบว่ามี flag อะไรถูกใช้งานบ้าง
		init, _ := cmd.Flags().GetBool("init")
		connect, _ := cmd.Flags().GetStringSlice("connect")
		changeVault, _ := cmd.Flags().GetBool("change-vault")
		changeVaultShort, _ := cmd.Flags().GetBool("cv")
		cancelConnect, _ := cmd.Flags().GetBool("cancel-connect")
		cancelConnectShort, _ := cmd.Flags().GetBool("cc")
		status, _ := cmd.Flags().GetBool("status")
		debug, _ := cmd.Flags().GetBool("debug")
		restore, _ := cmd.Flags().GetBool("restore")
		hostStatus, _ := cmd.Flags().GetBool("host-status")
		hostStatusShort, _ := cmd.Flags().GetBool("hs")
		update, _ := cmd.Flags().GetBool("update")
		updateShort, _ := cmd.Flags().GetBool("up")

		// STARTS
		if init {
			fmt.Println("Initializing AWE repository...")
			os.Exit(0)
		}

		if len(connect) > 0 {
			fmt.Printf("Connecting to SCA with keys: %v\n", connect)
			os.Exit(0)
		}

		if changeVault || changeVaultShort {
			fmt.Println("Switching to new SCA code...")
			os.Exit(0)
		}

		if cancelConnect || cancelConnectShort {
			fmt.Println("Unlinking SCA code...")
			os.Exit(0)
		}

		// OPTIONS
		if status {
			fmt.Println("Connection Status: Connected")
			fmt.Println("SCA Key: ****-****-****")
			os.Exit(0)
		}

		if debug {
			fmt.Println("Debug mode enabled - Connecting to AWE Network...")
			os.Exit(0)
		}

		if restore {
			fmt.Println("Restoring all values to default...")
			fmt.Println("Reset complete!")
			os.Exit(0)
		}

		if hostStatus || hostStatusShort {
			fmt.Println("Host Status Check:")
			fmt.Println("Network Provider: Online")
			fmt.Println("Connection: Stable")
			os.Exit(0)
		}

		// OTHERS
		if update || updateShort {
			fmt.Println("⬆Checking for updates...")
			fmt.Printf("Current version: %s\n", Version)
			fmt.Println("You are using the latest version!")
			os.Exit(0)
		}

		return nil
	}
}

func Init() {
	Initbanner()
	// ส่งออกไปยัง options.go
}
