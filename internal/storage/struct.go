package storage

type Banner struct {
	TagIds    []int
	FeatureID int
	Content   []byte
	IsActive  bool
}
