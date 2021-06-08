package uuid

import "github.com/google/uuid"

type UUID [16]byte

func (u UUID) String() string {
	impl := uuid.UUID(u)
	return impl.String()
}

func FromString(input string) (u UUID, err error) {
	impl, err := uuid.Parse(input)
	if err != nil {
		return u, err
	}
	u = UUID(impl)
	return
}

func Generate() UUID {
	impl, _ := uuid.NewUUID()
	return UUID(impl)
}
