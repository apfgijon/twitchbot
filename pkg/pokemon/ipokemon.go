package pokemon

type PokeInfo interface {
	PokeMoves(pokemon string) map[int]string
	PokeMovesFormatted(pokemon string) string
	Types(pokemon string) string
	TypeTable(typo string) string
	TypeTableFrom(typo string) string
	TypeTableTo(typo string) string
	CaptureRate(pokemon string) int
	Stats(pokemon string) string
	PokeEvos(pokemon string) string
	PokeRandom() string
}
