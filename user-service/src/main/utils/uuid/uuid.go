package uuid

import "github.com/nu7hatch/gouuid"

func V4() string {
	u, _ := uuid.NewV4()
	return u.String()
}
