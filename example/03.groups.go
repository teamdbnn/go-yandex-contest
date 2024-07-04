package yacontest

import (
	"context"
	"fmt"
)

func main() {
	c := NewClient("token")
	ans, err := c.NewCreateGroupForCompetition().Name("asdasdsdds").Do(context.Background())
	fmt.Println(ans, err)
	var a int
	fmt.Scan(&a)
}
