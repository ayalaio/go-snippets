package flyweight

import "time"

const (
	TEAM_A = iota
	TEAM_B
)

type Player struct {
	Name string
}

type Team struct {
	ID      string
	Name    string
	Players []Player
}

type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}

type Match struct {
	Date          time.Time
	VisitorTeamID uint64
	HostTeamID    uint64
	VisitorScore  uint64
	HostScore     uint64
}

type teamFlyweightFactory struct {
	createdTeams map[int]*Team
}

func NewFlyweightFactory() teamFlyweightFactory {
	return teamFlyweightFactory{make(map[int]*Team, 0)}
}

func getTeamFactory(teamName int) *Team {
	switch teamName {
	case TEAM_A:
		return &Team{
			ID:   "A",
			Name: "TEAM_A",
			Players: []Player{
				Player{"Juan Perez"},
			},
		}
	case TEAM_B:
		return &Team{
			ID:   "B",
			Name: "TEAM_B",
			Players: []Player{
				Player{"Nicolai Dominique"},
			},
		}
	}
	return nil
}

func (f *teamFlyweightFactory) GetTeam(teamName int) *Team {
	team := f.createdTeams[teamName]
	if team == nil {
		team = getTeamFactory(teamName)
		f.createdTeams[teamName] = team
	}
	return team
}

func (f *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(f.createdTeams)
}
