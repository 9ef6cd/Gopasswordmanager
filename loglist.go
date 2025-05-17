package main

import (
	"time"
	"fmt"
	"errors"
	"os"	
	"strconv"

	"github.com/aquasecurity/table"	
)

type Login struct {
	Website string
	Username string
	Password string
	Created time.Time
	LastChanged string
}

type Logins []Login

func (logins *Logins) add(website string, username string, password string, lastChanged string) {
	login := Login{
		Website: website,
		Username: username,
		Password: password,
		Created: time.Now(),
		LastChanged: lastChanged,	
	}	

	*logins = append(*logins, login)
}

func (logins *Logins) validateIndex(index int) error {
	if index < 0 || index >= len(*logins) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (logins *Logins) delete(index int) error {
	l := *logins

	if err := l.validateIndex(index); err != nil {
		return err
	}

	*logins = append(l[:index], l[index+1:]...)

	return nil

}

func (logins *Logins) edit(index int, website string, username string, password string, lastChanged string) error {
	l := *logins

	if err := l.validateIndex(index); err != nil {
		return err
	}

	l[index].Website = website
	l[index].Username = username
	l[index].Password = password
	l[index].LastChanged = lastChanged	
	return nil

}

func  (logins *Logins) Print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#","Website","Username","Password","Date Created","Last Changed")
	for index, l := range *logins {
		
		table.AddRow(strconv.Itoa(index), l.Website, l.Username,l.Password, l.Created.Format(time.RFC1123), l.LastChanged)
	}
	table.Render()
}
