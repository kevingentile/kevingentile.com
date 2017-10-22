package web

//Data is used to pass various child structures from a handler for templating
type Data map[string]interface{}

type Page struct {
	Title string
	Class string
}

//Carousel contains the resources to implement a reel carousel.
type Carousel struct {
	Articles []ReelArticle
}

//ReelArticle contains the structures used to create a carousel article.
type ReelArticle struct {
	Title     string
	Body      string
	ImageLink string
	PicPath   string
}

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

type PageHeading struct {
	Header string
	Body   string
	Button bool
}

//Nav contains the structure to implement a navigation bar.
type Nav struct {
	Buttons []NavButton
}

//NavButton contains the structure to implement a button in the navigation bar.
type NavButton struct {
	Name  string
	Link  string
	Class string
}

//Script holds the path to a script
type Script string

type Paragraph string

type PageArticle struct {
	Heading    string
	Paragraphs []Paragraph
}

type PageMain struct {
	Heading string
	Body    string
}
