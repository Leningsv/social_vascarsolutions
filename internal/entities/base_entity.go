package entities

func MigrateEntities() {
	MigrateUser()
	MigrateReport()
	MigratePost()
	MigrateComment()
}
