/*
 * @Description  : Core library of WebSocket Server
 * @Author       : Paddy Zhang
 * @Date         : 2021-08-05 13:46:01
 * @LastEditors  : Paddy Zhang
 * @LastEditTime : 2021-08-05 16:29:58
 * @FilePath     : \src\svr.go
 */
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func StartServer() {
	http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			// handle error
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " > ", "Srv start failed:", err)
		}
		go func() {
			defer conn.Close()

			for {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					// handle error
					// fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " > ", "Srv read msg failed:", err)
				}
				if string(msg) != "" {
					fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " > ", "Srv receive msg:", string(msg))
					err = wsutil.WriteServerMessage(conn, op, msg)
					if err != nil {
						// handle error
						// fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " > ", "Srv write msg failed:", err)
					}
					fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " > ", "Srv sent msg successful.")
				}
			}
		}()
	}))
}
