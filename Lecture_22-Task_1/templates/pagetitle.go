package pagetitle

type Title struct {
	PageTitle string
}

func NewTitle() *Title {
	return &Title{PageTitle: "Top Stories"}
}
