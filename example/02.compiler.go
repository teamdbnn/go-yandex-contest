package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
)

func main() {
	c := NewClient("token")
	ans, err := c.NewGetCompilersList().Do(context.Background())
	jsonData, err := json.Marshal(*ans)
	fmt.Println(string(jsonData), err)
	var a int
	fmt.Scan(&a)
}
