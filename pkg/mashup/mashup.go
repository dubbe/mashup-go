package mashup

type Artist struct {
	MBID        string
	Description string
	Name        string
	Albums      []Album
}

type Album struct {
	ID    string
	Title string
	Image string
}
