package ui

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func SelectItem(items []string, listItems bool) string {
	mp := make(map[int]string)
	for idx, item := range items {
		mp[idx] = item
		if listItems {
			fmt.Printf("\t%3d : %s\n", idx, item)
		}
	}
	fmt.Printf("Select: ")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	line = strings.TrimSpace(line)

	// string to int
	i, err := strconv.Atoi(line)
	if err != nil {
		fmt.Println("ERROR: Invalid selection...")
		return SelectItem(items, false)
	}

	item, found := mp[i]
	if !found {
		fmt.Println("ERROR: Invalid selection")
		return SelectItem(items, false)
	}

	fmt.Printf("selected: %s\n", item)

	return item
}
