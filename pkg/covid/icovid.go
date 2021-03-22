package covid

type CovidInfo interface {
	Build()
	GetCovidCasesForProvince(province string) (int, int)
	GetCovidCasesSpain() (int, int)
	FormatName(string) string
}
