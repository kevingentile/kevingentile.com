package web

type ContactBar struct {
	Header   string
	Body     string
	Contacts []ContactButton
}

type ContactButton struct {
	Label string
	Link  string
	Icon  string
}
