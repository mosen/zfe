package database

type Repository interface {
	Find(options ...func(*FindOptions) error) ([]interface{}, error)
}

type FindOptions struct {
	// Limit results returned to this number
	Limit uint64

	// Include relationships specified by name
	Include []string
}

func Limit(count uint64) func(*FindOptions) error {
	return func(opts *FindOptions) error {
		if count > 0 {
			opts.Limit = count
		} else {
			opts.Limit = 100
		}
		return nil
	}
}
