package main
import "api-reflist-migration-service/service"

var s service.RefListService

func main() {
	s.Migrate()
}
