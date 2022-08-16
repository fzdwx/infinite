package components

// Status About the state of the Component
type Status int

const (
	// Focus only use Input
	Focus Status = iota
	// Blur only use Input
	Blur
	// Quit user interrupt, should kill program
	Quit
	// Finish indicates that a component has done its job
	Finish
	// Normal ignore it
	Normal
)
