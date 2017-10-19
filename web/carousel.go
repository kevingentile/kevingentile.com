package web

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
