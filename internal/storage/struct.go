package storage

type Banner struct {
	TagIds    []int
	FeatureID int
	Content   Content
	IsActive  bool
}

type Content struct {
	Title string
	Text  string
	URL   string
}
