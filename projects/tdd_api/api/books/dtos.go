package books

type bookContentData struct {
	Isbn          string `json:"isbn,omitempty"`
	Title         string `json:"title,omitempty"`
	Image         string `json:"image,omitempty"`
	Genre         string `json:"genre,omitempty"`
	YearPublished int    `json:"year_published,omitempty"`
}
