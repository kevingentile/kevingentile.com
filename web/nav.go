package web

//Nav contains the structure to implement a navigation bar.
type Nav struct {
	Buttons []NavButton
}

//NavButton contains the struture to implement a button in the navigation bar.
type NavButton struct {
	Name  string
	Link  string
	Class string
}
