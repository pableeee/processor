package repository

// The repository packaget, provides the basis of an entity repository, to avoid boilerplate code.

// Repository abstracts the actual underlyin db infra, into a client interface.
type Repository interface {
	Get(id string, i interface{}) error
	Save(id string, i interface{}) error
	Update(id string, i interface{}) error
}

type repository struct {
}

func (r repository) Get(id string, i interface{}) error {
	return nil
}

func (r repository) Save(id string, i interface{}) error {
	return nil
}

func (r repository) Update(id string, i interface{}) error {
	return nil
}
