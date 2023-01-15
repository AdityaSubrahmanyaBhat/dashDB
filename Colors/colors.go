package colors

func GetColor(color string) string{
	switch(color){
	case "Red": return "\033[31m"
	case "Green": return "\033[32m"
	case "Yellow": return "\033[33m"
	case "Blue": return "\033[34m"
	case "White": return "\033[37m"
	default: return "\033[0m"
	}
}