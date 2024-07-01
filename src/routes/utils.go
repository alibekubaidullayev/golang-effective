package routes

func RouteMaker(method string, parentRoute string, address string) string {
	result := method
	result += " /" + parentRoute
	if address != "" {
		result += address
	}
	return result
}
