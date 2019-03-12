package flyweight

import "testing"

func TestTeamFlyweightFactory_GetTeam(t *testing.T) {
	factory := NewFlyweightFactory()

	teamA1 := factory.GetTeam(TEAM_A)
	if teamA1 == nil {
		t.Error("Te pointer to the TEAM_A was nil")
	}

	teamA2 := factory.GetTeam(TEAM_A)
	if teamA1 == nil {
		t.Error("Te pointer to the TEAM_A was nil")
	}

	if teamA1 != teamA2 {
		t.Error("Point to TeamA should be the same")
	}

	if factory.GetNumberOfObjects() != 1 {
		t.Errorf("Number of objects is not 1: %d\n", factory.GetNumberOfObjects())
	}
}

func TestHighVolume(t *testing.T) {
	factory := NewFlyweightFactory()

	teams := make([]*Team, 50000*2)

	for i := 0; i < 50000; i++ {
		teams[i] = factory.GetTeam(TEAM_A)
	}
	for i := 0; i < 50000; i++ {
		teams[i+50000] = factory.GetTeam(TEAM_B)
	}

	if teams[50000-1] == teams[50000] {
		t.Error("TeamA should be different than TeamB")
	}

	if teams[0] != teams[1] {
		t.Error("TeamA shuld always be teamA")
	}

	if teams[50000] != teams[50000+1] {
		t.Error("TeamB shuld always be teamB")
	}

	if factory.GetNumberOfObjects() != 2 {
		t.Error("Only teamA and teamB should exist")
	}
}
