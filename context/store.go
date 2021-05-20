package main

import "context"
import "log"

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "silly", "work")

	value := ctx.Value("silly")
	// 返り値はinterface{}なので値のキャストが必要
	str, ok := value.(string)
	if !ok {
		log.Println("value not found")
	}

	log.Printf("value found: '%s'", str)

	// 上の省略形がこれ
	str, ok = ctx.Value("silly").(string)
	log.Printf("re value found: '%s'", str)
}
