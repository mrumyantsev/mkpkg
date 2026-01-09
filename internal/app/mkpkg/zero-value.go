package mkpkg

func zeroValue(typ string) string {
	switch typ {
	case "string":
		return "\"\""
	case "bool":
		return "false"
	case "int8":
		fallthrough
	case "uint8":
		fallthrough
	case "byte":
		fallthrough
	case "int16":
		fallthrough
	case "uint16":
		fallthrough
	case "int32":
		fallthrough
	case "rune":
		fallthrough
	case "uint32":
		fallthrough
	case "int64":
		fallthrough
	case "uint64":
		fallthrough
	case "int":
		fallthrough
	case "uint":
		fallthrough
	case "uintprt":
		fallthrough
	case "float32":
		fallthrough
	case "float64":
		fallthrough
	case "complex64":
		fallthrough
	case "complex128":
		return "0"
	case "error":
		return "nil"
	default:
		typLength := len(typ)

		if typLength >= 1 && typ[0] == '*' ||
			typLength >= 2 && typ[:2] == "[]" ||
			typLength >= 4 && typ[:4] == "map[" {
			return "nil"
		}

		return typ + "{}"
	}
}
