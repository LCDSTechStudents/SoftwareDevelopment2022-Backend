package authorize

import "SoftwareDevelopment-Backend/server/internalsvc"

const (
	NAME = "AUTH"
)

type IAuthorizer interface {
	internalsvc.Internal
}

func GetName() string {
	return NAME
}
