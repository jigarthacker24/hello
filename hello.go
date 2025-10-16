package hello

func Say(name string, exited bool) string {
	if name == "" {
		name = "world"
	}

	greeting := "."
	if exited {
		greeting = "!!!"
	}

	return "Hello, " + name + greeting
}

func SayWithPrefix(prefix, name string, exited bool) string {
	if prefix == "" {
		prefix = "Hello"
	}
	if name == "" {
		name = "world"
	}

	greeting := "."
	if exited {
		greeting = "!!!"
	}

	return prefix + ", " + name + greeting
}
