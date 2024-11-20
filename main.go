package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/agungcandra/durian/cache"
)

func main() {
	m := cache.NewMemoryCache()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		command := scanner.Text()
		cmds := strings.Split(command, " ")

		switch cmds[0] {
		case "put":
			m.Put(cmds[1], cmds[2:]...)
		case "get":
			fmt.Println(m.Get(cmds[1]))
		case "delete":
			m.Delete(cmds[1])
		case "search":
			fmt.Println(m.Search(cmds[1], cmds[2]))
		case "keys":
			fmt.Println(m.Key())
		case "exit":
			os.Exit(0)
		}
	}
}
