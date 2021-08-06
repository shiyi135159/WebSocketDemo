/*
 * @Description  : Entrance of WebSocket Server
 * @Author       : Paddy Zhang
 * @Date         : 2021-08-05 13:46:01
 * @LastEditors  : Paddy Zhang
 * @LastEditTime : 2021-08-05 14:07:47
 * @FilePath     : \src\launcher.go
 */
package main

import (
	"fmt"
	"time"

	"WebSocketServer/tool"
)

func main() {
	const VER = "1.0"
	consoleTitle := fmt.Sprintf("WebSocket Server v%s is running - Don't close this window!", VER)
	tool.SetConsoleTitle(consoleTitle)
	go StartServer()
	for {
		time.Sleep(time.Minute)
	}
}
