package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.bin")
	if err != nil {
		// no such file やら permission errorやら
		// このハンドリングを省略するとロクなことがない
		log.Fatal(err)
	}
	// このタイミングでクローズのdeferを書いておく
	defer file.Close()

	// すべてを[]byteで読み出す荒業
	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(b)

	/// open からReadallまでの流れを全部やってくれるすごいやつ
	content, err := ioutil.ReadFile("test.bin")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(content)
}
