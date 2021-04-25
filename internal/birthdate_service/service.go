package birthdate_service

import (
	"fmt"
	"log"
)

type Service struct {
	Repo Repo
}

type Repo interface {
	InsertUserDateWithChannel(int64, string, string) error
}

func (s *Service) SetBirthdate(groupTgID int64, login string, month int, day int) error {

	// format date to string
	dateString := fmt.Sprintf("2020-%d-%d", month, day)

	log.Printf("Setting birthday for %s, group %d, date %s", login, groupTgID, dateString)

	// set birthdate by repo method
	err := s.Repo.InsertUserDateWithChannel(groupTgID, login, dateString)
	if err != nil {
		log.Printf("Failed to set birthdate %s", dateString)
		return err
	}

	log.Printf("Birthdate %s is set", dateString)

	return nil
}

func InitService(repo Repo) *Service {
	return &Service{Repo: repo}
}
