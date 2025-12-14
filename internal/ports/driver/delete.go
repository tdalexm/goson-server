package portsdriver

type DeleteService interface {
	Execute(collectionType, id string) (string, error)
}
