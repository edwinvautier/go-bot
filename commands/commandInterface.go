package commands

// Command is the interface used to handle all our commands (meteo, music, ai...)
type Command interface {
	Execute()
}