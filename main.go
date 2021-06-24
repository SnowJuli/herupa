package main

import (
		"fmt"
		"strings"
		
)

func main() {
	var choice string

	fmt.Println(`
Herupa - A Hackintosh Manager
-----------------------------
- Backup your EFI (Backup)
- Fetch Kexts (Fetch)
- Gather new Opencorepkg (Opencore)
- Transfer files between Backup and EFI (Transfer)
-----------------------------
Enter the Words in Brackets:`)

	fmt.Scanf("%s", &choice)
	var choice_lwr = strings.ToLower(choice)

	switch {
	case choice_lwr == "backup":
		return
	case choice_lwr == "fetch":
		return
	case choice_lwr == "opencore":
		return
	case choice_lwr == "transfer":
		return
	default:
		fmt.Println("Please enter the right command.")
	}
}
