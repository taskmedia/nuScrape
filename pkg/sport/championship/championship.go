package championship

// Allowed championship for BHV
var championshipAbb = [...]string{
	"BHV",
	"UF",
	"OF",
	"MF",
	"OS",
	"SW",
	"AB",
	"AV",
	"OB",
}

// ValidateChampionshipAbb checks if the given championship abbreviation is valid
func ValidateChampionshipAbb(c string) bool {
	for _, ca := range championshipAbb {
		if ca == c {
			return true
		}
	}
	return false
}
