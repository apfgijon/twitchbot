package filesystem

import (
	"encoding/json"
	"io/ioutil"
)

type FileProviderv1 struct {
}

func NewFileProvider() FileProvider {
	return &FileProviderv1{}
}

func (this *FileProviderv1) GetCounterCommand(command string) int {
	read, _ := ioutil.ReadFile("counter.json")
	var counterCommands map[string]int

	json.Unmarshal(read, &counterCommands)

	counterCommands[command]++

	write, _ := json.MarshalIndent(counterCommands, "", "	")

	ioutil.WriteFile("counter.json", write, 0644)
	return counterCommands[command]
}
