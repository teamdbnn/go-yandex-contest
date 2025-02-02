package yacontest

import (
	"context"
	"fmt"
)

func main() {
	c := NewClient("token")
	err := c.NewDeleteGroupFromContest().ContestID(63561).GroupID(3435338).Do(context.Background())

	fmt.Println(err)
}
