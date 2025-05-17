package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"os"
	"time"
)
type CmdFlags struct {
	Add string
	Del int
	Edit string
	List bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "Add", "",  "Adds new entry format: -Add website:username")
	flag.StringVar(&cf.Add, "add", "",  "Adds new entry format: -Add website:username")
	flag.StringVar(&cf.Edit, "Edit", "", "Edits a selected entry by index generates a new passwod on edit leave website/username empty to keep values format: -Edit Index:Website:Username")
	flag.IntVar(&cf.Del, "Del", -1, "Deletes entry on index selected Format: -Del Index")
	flag.BoolVar(&cf.List, "List", false, "List all entries")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(logins *Logins) {
	switch {
	case cf.List:
		logins.Print()
	case cf.Add != "":
		parts := strings.SplitN(cf.Add, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error: Add must be formated website:username")
			os.Exit(1)
		}
		
		website := parts[0]
		username := parts[1] 

		password := Passgen()
		lastChanged := "N/A"	

		logins.add(website, username, password, lastChanged)
	
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 4)
		l := *logins 

		if len(parts) != 3 {
			fmt.Println("error invalid format")
			os.Exit(1)
		}
		
		index, err := strconv.Atoi(parts[0]) 
		website := parts[1]
		username := parts[2]

		if parts[1] == "" {
		website = l[index].Website
		}
		
		if parts[2] == "" {
		username = l[index].Username
		}

		if err != nil {
			fmt.Println("Error: invald index please try again")
			os.Exit(1)
		}

		password := Passgen()
		lastChanged := time.Now().Format(time.RFC1123)
		
		logins.edit(index, website, username, password, lastChanged)

	case cf.Del != -1:
		var input string

		fmt.Printf("are you sure you wanna delete this entry Y/N :")
		fmt.Scan(&input)
		
		if input == "y" || input == "Y" {
		logins.delete(cf.Del)
		fmt.Println("Entry deleted")
		}

		if input == "n" || input == "N" {
		os.Exit(1)
		fmt.Println("Operation canceled")
		}

	default:
		fmt.Println("invalid command")
	}
}

