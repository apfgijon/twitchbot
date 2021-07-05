package covid

import "strings"

func (cI *CovidApiImpl) FormatName(province string) string {
	province = strings.ToLower(province)
	switch province {
	case "andalucia", "andalucía":
		return "andalusia"
	case "aragon", "aragón":
		return "aragon"
	case "asturias", "asturies":
		return "asturias"
	case "baleares", "islas baleares":
		return "baleares"
	case "valencia", "c. valencia", "comunidad valenciana":
		return "c. valenciana"
	case "canarias", "islas canarias":
		return "canarias"
	case "cantabria":
		return "cantabria"
	case "castilla la mancha":
		return "castilla - la mancha"
	case "castilla y leon", "castilla y león":
		return "castilla y leon"
	case "cataluña", "catalunya":
		return "catalonia"
	case "ceuta":
		return "ceuta"
	case "extremadura":
		return "extremadura"
	case "galicia":
		return "galicia"
	case "la rioja":
		return "la rioja"
	case "madrid":
		return "madrid"
	case "melilla":
		return "melilla"
	case "murcia":
		return "murcia"
	case "navarra":
		return "navarra"
	case "pais vasco", "país vasco":
		return "pais vasco"
	}
	return ""
}
