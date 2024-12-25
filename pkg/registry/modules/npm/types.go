package npm

// Packages

type RegistryPackage struct {
	Id				string										`json:"_id"`
	Name			string										`json:"name"`
	Maintainers		[]RegistryPackageAuthor						`json:"maintainers"`
	DistTags		map[string]string							`json:"dist-tags,omitempty"`
	Versions		map[string]RegistryPackageVersion			`json:"versions"`
	Dependencies 	map[string]RegistryPackageDependencyInfo	`json:"dependencies"`
	NPMUsers		[]string									`json:"npmUsers"`
}

type RegistryPackageDependencyInfo struct {
	Version	string
	Exists	bool
}

type RegistryPackageVersion struct {
	Name			string						`json:"name"`
	Repository		RegistryPackageRepository	`json:"repository"`
	Dependencies 	map[string]string			`json:"dependencies"`
	DevDependencies map[string]string			`json:"devDependencies,omitempty"`
	NPMUser			RegistryPackageAuthor		`json:"_npmUser"`
}

type RegistryPackageRepository struct {
	URL		string	`json:"url"`
}

type RegistryPackageAuthor struct {
	Name	string	`json:"name"`
	Email	string	`json:"email"`
}

// Package JSON File

type NPMPackageJSON struct {
	Name         string            `json:"name"`
	Version      string            `json:"version"`
	Dependencies map[string]string `json:"dependencies"`

	Main            string            `json:"main,omitempty"`
	Author          string            `json:"author,omitempty"`
	License         string            `json:"license,omitempty"`
	Homepage        string            `json:"homepage,omitempty"`
	DevDependencies map[string]string `json:"devDependencies,omitempty"`
	Scripts         map[string]string `json:"scripts,omitempty"`
	Description     string            `json:"description,omitempty"`
	Repository      interface{}       `json:"repository,omitempty"`
	Bugs            interface{}       `json:"bugs,omitempty"`
	Keywords        []string          `json:"keywords,omitempty"`
}
