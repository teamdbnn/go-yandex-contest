package yacontest

import (
	"context"
	"fmt"
)

func main() {
	c := NewClient("token")
	err := c.NewStartContestForParticipant().ContestID(63561).ParticipantID(118779979).Do(context.Background())

	fmt.Println(err)
}
