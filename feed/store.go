package feed

// FeedStore is the db interface
type FeedStore interface {
	// GetPosition retrieve the last position from the db
	GetPosition(table string) uint64
	// SetPosition save the new feed position in the db
	SetPosition(table string, last uint64) error
	// Save the feed in the db
	Save(last interface{}) error
}
