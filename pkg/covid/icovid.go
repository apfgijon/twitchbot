package covid

type CovidInfo interface {
	GetCovidCasesForProvince(province string) (int, int)
	GetCovidCasesSpain() (int, int)
	FormatName(string) string
}
