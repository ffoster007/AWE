package connect

import (
    "fmt"
    "os"
)



func getAccessKey() string {
    // รับ SCA key จาก cli flag ของ User 
    accessKeys, _ := RootCmd.Flags().GetStringSlice("connect")
    if len(accessKeys) > 0 {
        return accessKeys[0] // คืนค่า SCA key แรกที่ได้รับ
    }
    return "" // คืนค่าว่างถ้าไม่มี SCA key
}
        if hostStatus || hostStatusShort {
            fmt.Println("Host Status: All systems operational.")
            os.Exit(0)
        }
