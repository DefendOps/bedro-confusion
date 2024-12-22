package npm

type RegistryPackage struct {
	Id		string					`json:"_id"`
	Name	string					`json:"name"`
	Author	RegistryPackageAuthor	`json:"author"`
}

type RegistryPackageAuthor struct {
	Name	string	`json:"name"`
	Email	string	`json:"email"`
}