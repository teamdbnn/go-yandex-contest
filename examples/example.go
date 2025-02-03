package main

import (
	"context"
	"fmt"
	yacontest "github.com/teamdbnn/go-yandex-contest/v2"
)

func main() {
	c := yacontest.NewClient("token")

	// Get participants
	participants, err := c.NewGetParticipantsOfCompetition().CompetitionID(3696).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(participants)

	compilers, err := c.NewGetCompilersList().Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(compilers)

	group, err := c.NewCreateGroupForCompetition().Name("asdasdsdds").Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(group)

	if err = c.NewDeleteGroupFromContest().ContestID(63561).GroupID(*group.ID).Do(context.Background()); err != nil {
		panic(err)
	}
}
