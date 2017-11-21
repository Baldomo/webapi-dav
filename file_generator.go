package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args[3:]
	numPtr := flag.Int("num", 80, "Indica quanti file casuali saranno generati")
	for _, arg := range args {
		if abs, err := filepath.Abs(arg); err != nil {
			continue
		} else {
			arg = abs
		}
		for i := 0; i <= *numPtr; i++ {
			file, err := os.Create(filepath.Join(arg, fmt.Sprint(i, ".txt")))
			if err != nil {
				log.Fatal(err)
			}

			if err := file.Truncate(randInRange(4e5, 4e6)); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func randInRange(min int, max int) int64 {
	return int64(rand.Intn(max-min) + min)
}
