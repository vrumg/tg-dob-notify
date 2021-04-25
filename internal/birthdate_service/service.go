package birthdate_service

// Service main birthday service class
type Service struct {
	Repo Repo
}

// Repo repository interface required by Service class
type Repo interface {
	InsertUserDateWithChannel(int64, string, string) error
}

// InitService initialize and populate new Service object
func InitService(repo Repo) *Service {
	return &Service{Repo: repo}
}
