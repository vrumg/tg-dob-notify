package birthdate_service

import (
	"fmt"
	"github.com/rs/zerolog/log"
)

// SaveUserBirthdate set or update user birthdate
func (s *Service) SaveUserBirthdate(group int64, login string, month int, day int) error {
	log.Debug().
		Str("login", login).
		Int64("group", group).
		Msg("setting new birthday")

	// format date to string
	dateString := fmt.Sprintf("2020-%d-%d", month, day)

	err := s.Repo.InsertUserDateWithChannel(group, login, dateString)
	if err != nil {
		log.Warn().
			Str("login", login).
			Msg("failed to set birthday")

		return err
	}

	log.Debug().
		Str("login", login).
		Msg("setted birthday")

	return nil
}
