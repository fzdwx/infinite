package components

// Status About the state of the Component
type Status int

const (
	// Focus only use Input
	Focus Status = iota
	// Blur only use Input
	Blur
	// Quit component
	Quit
	// Normal ignore it
	Normal
)
