package interfaces

type Register = func(cb func(m any, aliases []string, examples [][2]string))
