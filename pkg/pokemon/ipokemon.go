package pokemon

type PokeInfo interface {
	Build()
	PokeMoves(pokemon string, gmaeVersion string) map[int]string
	PokeMovesFormatted(pokemon string, gmaeVersion string) string
	Types(pokemon string) string
	TypeTable(typo string) string
	TypeTableFrom(typo string) string
	TypeTableTo(typo string) string
	CaptureRate(pokemon string) int
	Stats(pokemon string) string
	PokeEvos(pokemon string) string
}
