package main

import (
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/gofrs/uuid/v5"
)

func main() {
	var sb strings.Builder
	_, _ = sb.WriteString(`INSERT INTO public.users (id,"name",age) VALUES`)
	_ = sb.WriteByte('\n')
	for i := 0; i < 3000; i++ {
		if i != 0 {
			_, _ = sb.WriteString(",\n")
		}
		id := uuid.Must(uuid.NewV4())
		_, _ = sb.WriteString(fmt.Sprintf("('%s','%s',%d)", id, "daniel", rand.IntN(30)))
	}
	_ = sb.WriteByte(';')
	fmt.Println(sb.String())
}
