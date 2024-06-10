package service

import (
	"log"
	"strconv"
)

func (s *Service) DeleteUser(idS string) error {
	id, err := strconv.Atoi(idS)
	if err != nil {
		log.Println("cannot convert id to int", err.Error())
		return err
	}
	return s.Repository.DeleteUser(int64(id))
}
