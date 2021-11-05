package handler

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/duy9911/Team/handler/logger"
	"github.com/duy9911/Team/handler/redis"
	"github.com/duy9911/Team/models"
)

type Id struct {
	Domain string
	Id     int
}

func CreateTeam(team models.Team) {
	validatedTeam, err := prepareTeam(team)
	if err != nil {
		logger.Logger("error prepare data team ", err)
		return
	}
	redis.Set(validatedTeam.ID, validatedTeam)
	logger.Logger("Create successfull team's information ", validatedTeam)
}

func ReturnTeams(domain string) {
	teams, err := redis.GetAll(domain)
	if err != nil {
		logger.Logger("error get all ", err)
		return
	}
	for _, v := range teams {
		logger.Logger("Team: ", v)
	}
}

func UpdateTeam(key string, team models.Team) {
	updateTeam, err := prepareUpdate(key, team)
	if err != nil {
		logger.Logger("error prepare update team", err)
		return
	}

	if err := redis.Set(key, updateTeam); err != nil {
		logger.Logger("error update", err)
		return
	}
	logger.Logger("Updated ", updateTeam)
}

func Deletestaff(key string) {
	_, err := redis.Get(key)
	if err != nil {
		logger.Logger("error key", errors.New("doesn't match any key"))
		return
	}

	if err := redis.Delete(key); err != nil {
		logger.Logger("error delete", err)
		return
	}
	logger.Logger("deleted ", key)
}

func validateTeam(team models.Team) error {

	if team.Name == " " {
		return errors.New("opp! team can not empty")
	}
	return nil
}

func prepareUpdate(key string, team models.Team) (models.Team, error) {
	updateTeam := models.Team{}

	if _, err := redis.Get(key); err != nil {
		return updateTeam, errors.New("doesn't match any key ")
	}
	err := validateTeam(team)
	if err != nil {
		return updateTeam, err
	}
	updateTeam = models.Team{
		ID:   key,
		Name: team.Name,
	}
	return updateTeam, err
}

func prepareTeam(s models.Team) (models.Team, error) {
	team := models.Team{}
	err := validateTeam(s)
	if err != nil {
		return s, err
	}
	domain := "team"
	nextKey, err := GenerateId(domain)
	if err != nil {
		return team, err
	}
	team = models.Team{
		ID:   nextKey,
		Name: s.Name,
	}
	return team, nil
}

// team-id: N+1
// team-id: N+1

// teams: Hashes
// teams: Hashes

func GenerateId(domain string) (string, error) {
	keyLasest := "lastId"
	idLatest, err := redis.Get(keyLasest)

	// check lastest key is empty or not
	if err == nil {
		idStruct := Id{}
		err := json.Unmarshal([]byte(idLatest), &idStruct)
		if err != nil {
			return domain, err
		}
		nextId := &Id{
			Domain: idStruct.Domain,
			Id:     idStruct.Id + 1,
		}
		redis.Set(keyLasest, nextId)
		concatenated := nextId.Domain + strconv.Itoa(nextId.Id)
		return concatenated, nil
	}

	newId := Id{
		Domain: domain,
		Id:     1,
	}
	redis.Set(keyLasest, newId)
	concatenated := newId.Domain + strconv.Itoa(newId.Id)
	return concatenated, nil
}
