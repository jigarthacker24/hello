package hello

func Say(name string) string {
	if name == "" {
		name = "world"
	}

	return "Hello, " + name + "!"
}

func SayWithPrefix(prefix, name string) string {
	if prefix == "" {
		prefix = "Hello"
	}
	if name == "" {
		name = "world"
	}
	return prefix + ", " + name + "!"
}
