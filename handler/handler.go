package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	models "github.com/duy9911/Team/model"
	"github.com/duy9911/Team/pkgs/logger"
	"github.com/duy9911/Team/pkgs/redis"
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
	redis.HashSet(validatedTeam.ID, validatedTeam)
	logger.Logger("Create successfull team's information ", validatedTeam)
}

func ReturnTeams() {
	teams, err := redis.HashGetAll()
	if err != nil {
		logger.Logger("error get all ", err)
		return
	}
	for _, v := range teams {
		logger.Logger("Team: ", v)
	}
}

func UpdateTeam(team_id string, team models.Team) {
	updateTeam, err := prepareUpdate(team_id, team)
	if err != nil {
		logger.Logger("error prepare update team", err)
		return
	}

	if err := redis.HashSet(team_id, updateTeam); err != nil {
		logger.Logger("error update", err)
		return
	}
	logger.Logger("Updated ", updateTeam)
}

func DeleteTeams(team_id string) {
	result, err := redis.HDel(team_id)
	if err != nil {
		logger.Logger("error delete", errors.New("error delete "+team_id))
		return
	}
	if result == 0 {
		logger.Logger("error delete", errors.New("wrong key "+team_id))
	}
	logger.Logger("deleted ", team_id)
}

//add staff
func AddStaffsToTeam(team_id string, staffs []string) {
	fmt.Println(staffs)
	fmt.Println(staffs[1])
	team, err := prepareStaffsToTeam(team_id, staffs)
	if err != nil {
		logger.Logger("error prepare", err)
	}
	if err := redis.HashSet(team_id, team); err != nil {
		logger.Logger("error add staffs to team ", err)
	}
	logger.Logger("success add staffs to team", team_id)
}

func prepareStaffsToTeam(team_id string, staffs []string) (models.Team, error) {
	teamStruct := models.Team{}
	team, err := redis.HashGet(team_id)
	if err != nil {
		return teamStruct, err
	}
	errUnmar := json.Unmarshal([]byte(team), &teamStruct)
	if errUnmar != nil {
		return teamStruct, err
	}
	teamUpdated := models.Team{
		ID:     team_id,
		Name:   teamStruct.Name,
		Staffs: staffs,
	}
	return teamUpdated, nil
}

func prepareUpdate(team_id string, team models.Team) (models.Team, error) {
	updateTeam := models.Team{}

	if _, err := redis.HashGet(team_id); err != nil {
		return updateTeam, errors.New("doesn't match any key ")
	}
	err := validateTeam(team)
	if err != nil {
		return updateTeam, err
	}
	updateTeam = models.Team{
		ID:   team_id,
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

func validateTeam(team models.Team) error {
	if team.Name == " " {
		return errors.New("opp! team can not empty")
	}
	return nil
}

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
