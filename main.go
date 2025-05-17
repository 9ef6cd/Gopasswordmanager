package main

func main() {
	logins := Logins{}
	storage := NewStorage[Logins]("Logins.json")
	storage.Load(&logins)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&logins)
	storage.Save(logins)
}

