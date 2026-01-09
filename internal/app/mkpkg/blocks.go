package mkpkg

import (
	"strings"

	"github.com/mrumyantsev/mkpkg/internal/logging"
)

func (a *App) addBlockPackage(sb *strings.Builder) {
	// package start
	sb.WriteString("package ")
	sb.WriteString(a.config.PackageName)
	sb.WriteString(a.config.Eol)
	// package end

	logging.Infof("created package '%s'\n", a.config.PackageName)
}

func (a *App) addBlockObject(sb *strings.Builder) {
	// object start
	sb.WriteString(a.config.Eol)
	sb.WriteString("type ")
	sb.WriteString(a.config.ObjectName)
	sb.WriteString(" ")
	sb.WriteString(a.config.ObjectType)
	sb.WriteString(" {")
	sb.WriteString(a.config.Eol)
	// interface methods start
	if a.config.ObjectType == "interface" {
		for i := range a.config.ObjectMethods {
			sb.WriteString("\t")
			sb.WriteString(a.config.ObjectMethods[i])
			sb.WriteString(a.config.Eol)

			logging.Infof("created method '%s'\n", a.config.ObjectMethods[i])
		}
	}
	// interface methods end
	sb.WriteString("}")
	sb.WriteString(a.config.Eol)
	// object end

	logging.Infof("created %s '%s'\n", a.config.ObjectType, a.config.ObjectName)

	if a.config.ObjectType == "struct" {
		// object constructor start
		sb.WriteString(a.config.Eol)
		sb.WriteString("func ")
		sb.WriteString(a.config.ObjectConstructor)
		sb.WriteString(" *")
		sb.WriteString(a.config.ObjectName)
		sb.WriteString(" {")
		sb.WriteString(a.config.Eol)
		sb.WriteString("\treturn &")
		sb.WriteString(a.config.ObjectName)
		sb.WriteString("{}")
		sb.WriteString(a.config.Eol)
		sb.WriteString("}")
		sb.WriteString(a.config.Eol)
		// object constructor end

		logging.Infof("created constructor '%s'\n", a.config.ObjectConstructor)

		// object methods start
		for i := range a.config.ObjectMethods {
			sb.WriteString(a.config.Eol)
			sb.WriteString("func (")
			sb.WriteString(a.config.ObjectReceiver)
			sb.WriteString(" *")
			sb.WriteString(a.config.ObjectName)
			sb.WriteString(") ")
			sb.WriteString(a.config.ObjectMethods[i])
			sb.WriteString(" {")
			sb.WriteString(a.config.Eol)
			// return values start
			idx := strings.IndexByte(a.config.ObjectMethods[i], ')') + 1
			if idx < len(a.config.ObjectMethods[i]) {
				sb.WriteString("\treturn ")
				typs := strings.Split(strings.Trim(strings.ReplaceAll(a.config.ObjectMethods[i][idx+1:], " ", ""), "()"), ",")
				typsLength := len(typs)
				sb.WriteString(zeroValue(typs[0]))
				for i := 1; i < typsLength; i++ {
					sb.WriteString(", ")
					sb.WriteString(zeroValue(typs[i]))
				}
				sb.WriteString(a.config.Eol)
			}
			// return values end
			sb.WriteString("}")
			sb.WriteString(a.config.Eol)

			logging.Infof("created method '%s'\n", a.config.ObjectMethods[i])
		}
		// object methods end
	}
}
