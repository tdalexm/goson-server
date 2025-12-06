package domain

type Repository interface {
	List(resource string) ([]Record, error)
	ListWithFilter(resource string, filter Filter) ([]Record, error)
	GetByID(resource, id string) (Record, error)
}
