package clone

//StartCloneProcess -
func StartCloneProcess() {
	svc := newService()
	svc.ValidateOrg()
	svc.ValidateDestination()
	svc.CloneReposToDestination()
}
