package yacontest

import (
	"context"
	"fmt"
)

func main() {
	c := NewClient("token")
	roles := []string{"ADMIN", "COACH"}
	err := c.NewUpdateGroupForContest().ContestID(63561).GroupID(3435331).Roles(roles).Do(context.Background())
	fmt.Println(err)
	var a int
	fmt.Scan(&a)
}
