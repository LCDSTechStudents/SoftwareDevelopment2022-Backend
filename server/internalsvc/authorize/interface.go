package authorize

import "SoftwareDevelopment-Backend/server/internalsvc"

type IAuthorizer interface {
	internalsvc.Internal
}

func GetName() string {
	return NAME
}
