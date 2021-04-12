package prov

type MessageProvider interface {
	Build()
	GetBingoCartonResponse() string
	GetPokemonRandomResponse(user string) string
	GetPokemonAtacksResponse(poke string) string
	GetPokemonEvolutionResponse(poke string) string
	GetPokemonTypesResponse(poke string) string
	GetPokemonCaptureRateResponse(poke string) string
	GetPokemonStatsResponse(poke string) string
	GetPokemonTypeTableResponse(typ string) string
	GetCovidStatsResponse(site string, user string) string
	GetUnderLevelResponse(command string) string
	GetBotellaResponse(getusers func() ([]string, error), u string) string
}
