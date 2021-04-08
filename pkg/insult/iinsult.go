package insult

type Insult interface {
	Build()
	Insult() string
}
