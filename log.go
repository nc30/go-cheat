package main

import (
	"errors"
	"log"
)

func main() {
	log.Println("spam")                     // 2021/01/19 18:47:35 spam
	log.Println("egg", "bacon", "and spam") // 2021/01/19 18:51:04 egg bacon and spam

	// Printfはfmtと同じ
	log.Printf("I'm a %s It's Ok", "Lumberjack") // 2021/01/19 19:53:54 I'm a Lumberjack It's Ok
	log.Printf("int: %d", 253)                   // 2021/01/19 19:51:39 int: 253
	log.Printf("hex: 0x%x", 253)                 // 2021/01/19 19:51:39 hex: 0xfd
	log.Printf("oct: 0o%o", 253)                 // 2021/01/19 19:51:39 oct: 0o375
	log.Printf("bin: 0b%b", 253)                 // 2021/01/19 19:51:39 bin: 0b11111101

	s := struct {
		ID   int
		Name string
	}{123, "Graham"}

	// 構造体のダンプ時に便利
	log.Printf("%+v", s) // 2021/01/19 19:50:00 {ID:123 Name:Graham}

	log.SetPrefix("[log] ")
	log.Println("プレフィックスをつける") // [log] 2021/01/19 18:50:07 プレフィックスをつける
	log.SetPrefix("")

	log.SetFlags(log.Flags() | log.LUTC)
	log.Println("時刻タイムゾーンはデフォルトを使用するが、フラグを追加することによりUTCにできる") // 2021/01/19 09:57:09 時刻タイムゾーンはデフォルトを使用するが、フラグを追加することによりUTCにできる

	log.SetFlags(0)
	log.Println("フラグをすべて外すと時間出力をオフにできる") // フラグを設定すると時間表示をオフにできる

	log.SetFlags(log.Ldate | log.Lmicroseconds)
	log.Println("マイクロ秒まで表記") // 2021/01/19 18:57:09.086480 マイクロ秒まで表記

	log.Fatal(errors.New("何かしらのエラー")) // エラー内容を出力してstatus code 1で終了する
}
