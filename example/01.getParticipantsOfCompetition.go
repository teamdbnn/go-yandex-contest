package yacontest

import (
	"context"
	"fmt"
)

func main() {
	c := NewClient("token")
	ans, err := c.NewGetParticipantsOfCompetition().CompetitionID(3696).Do(context.Background())
	fmt.Println(ans, err)
	var a int
	fmt.Scan(&a)
}
