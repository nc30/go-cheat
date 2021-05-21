// httpサーバー等でgraceful shutdownを行う場合や、
// 終了時に特別な処理をしたい場合はシグナルを監視する
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sig := make(chan os.Signal, 1)

	// SIGTERMとSIGINGをchanに転送する
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	// signalシグナル受信まで待つ
	s := <-sig
	log.Printf("signal: [%s] receive.\r\n", s)
}
