package hello

func Say(name string) string {
	if name == "" {
		name = "world"
	}

	return "Hello, " + name + "!"
}
