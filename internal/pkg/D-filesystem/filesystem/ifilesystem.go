package filesystem

type FileProvider interface {
	GetCounterCommand(command string) int
}
