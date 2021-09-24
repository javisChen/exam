package uuid

import uuid "github.com/satori/go.uuid"

func UUID() string {

	// or error handling
	v4 := uuid.NewV4()
	return v4.String()
}
