package npm

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/defendops/bedro-confuser/pkg/registry"
	"github.com/defendops/bedro-confuser/pkg/registry/payloads"
	"github.com/defendops/bedro-confuser/pkg/utils/requester"
	"github.com/defendops/bedro-confuser/pkg/utils/source"
	"github.com/defendops/bedro-confuser/pkg/utils/types"
)

type RegistryModule struct {
	GithubAPI *payloads.GithubAPI
	channel chan string
}

func init() {
	GithubAPI := payloads.NewGithubAPI()
	npmModule := RegistryModule{
		GithubAPI: GithubAPI,
		channel: make(chan string),
	}

	registry.RegisterRegistryModule(&npmModule)
}

func (n *RegistryModule) Run(src source.Source, scan_config types.Config, ctx *context.Context) error {
	// AdaptedSource is expected to be interface{} type
	// 
	// @PackageNameSource Scenarios
	// When we're looking up a Single Package from the source.PackageNameSource
	// We're expecting the adaptedSource interface to have the following keys ["Vulnerable", "Content"] and an empty Error
	// 
	// @URLSource Scenarios
	// Technically this module won't be executed if the src.Registry is anything other than "npm" meaning that this URL has to be a package.json or package-lock.json
	// Which is identified based on the SourceIdentifier function, therefore the expected adaptedsource is a map[string]interface{} containing ["notFoundPackages", "foundPackages"]
	adaptedSource, err := n.SourceAdapter(src)
	if err != nil{
		fmt.Println(err)
		return err
	}

	switch src.Type {
		case source.PackageNameSource:
			var isVulnerable bool
			var packageContent RegistryPackage
			var args map[string]interface{}

			k := reflect.ValueOf(adaptedSource)
			if k.Kind() == reflect.Map{
				for _, key := range k.MapKeys(){
					switch strings.ToLower(key.String()){
						case "vulnerable":
							isVInterface := k.MapIndex(key).Interface()
							
							// Type Casting
							isVulnerable, _ = isVInterface.(bool)
						case "content":
							pCInterface := k.MapIndex(key).Interface()
							
							// Type Casting
							packageContent, _ = pCInterface.(RegistryPackage)
					}
				}
			}

			if !isVulnerable{
				fmt.Printf("[i] %s Package was found\n", src.RawValue)
			}else{
				fmt.Printf("[i] %s Package was not found\n", src.RawValue)
				
			}

			packageContent.Dependencies[src.RawValue] = RegistryPackageDependencyInfo{
				Exists: !isVulnerable,
				Version: "99.99.99",
			}

			for dep, depInfo := range packageContent.Dependencies{
				if !depInfo.Exists && scan_config.CreatePackages{
					go func(dep string){
						args = map[string]interface{}{
							"packageName": dep,
						}

						n.CreatePackage(args)
					}(dep)

					<-n.channel
				}
			}

		case source.URLSource:
			packages_to_takeover := []string{}
			
			foundPkgs := make(map[string]string)
			notFoundPkgs := make(map[string]string)

			k := reflect.ValueOf(adaptedSource)
			if k.Kind() == reflect.Map {
				for _, key := range k.MapKeys() {
					switch strings.ToLower(key.String()) {
					case "found":
						fPkgInterface := reflect.ValueOf(k.MapIndex(key).Interface())
						if fPkgInterface.Kind() == reflect.Slice {
							foundPkgs = processSliceOfMaps(fPkgInterface)
						}
					case "not_found":
						nfPkgInterface := reflect.ValueOf(k.MapIndex(key).Interface())
						if nfPkgInterface.Kind() == reflect.Slice {
							notFoundPkgs = processSliceOfMaps(nfPkgInterface)
						}
					}
				}
			}

			for pkg := range notFoundPkgs{
				if !Contains(packages_to_takeover, pkg){
					packages_to_takeover = append(packages_to_takeover, pkg)
				}
			}
			
			for pkg := range foundPkgs{
				
				go func(pkg string) {
					tempSrc := source.Source{
						Type: source.PackageNameSource,
						RawValue: pkg,
						Registry: src.Registry,
					}

					n.Run(tempSrc, scan_config, ctx)
					n.channel<-pkg
				}(pkg)
				
				<-n.channel
			}

		default:
			fmt.Println(err)
	}

	return nil
}

func (n *RegistryModule) GonnaBeExecuted(registry string) bool {
	return registry == string(n.Registry())
}

func (n *RegistryModule) Name() string {
	return "NPM - Demo"
}

func (n *RegistryModule) Registry() registry.Registry {
	return "npm"
}

func (n* RegistryModule) LookupPackage(package_name string) ([]byte, error){
	request := requester.HTTPRequest{
		BaseURL: "https://registry.npmjs.org",
		Endpoint: fmt.Sprintf("/%s", package_name),
		Method: "GET",
		IsJson: false,
		Body: "",
		Headers: map[string]string{},
	}
	http := requester.HTTPRequester{}
	
	resp, err := http.PerformRequest(request)
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()

	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading URL response: %v", err)
	}

	return response, nil
}

func (n *RegistryModule) SourceAdapter(src source.Source) (interface{}, error) {
	switch src.Type{
		case source.PackageNameSource:
			isFound := true
			package_name := src.RawValue
			package_name = strings.ToLower(package_name)

			response, err := n.LookupPackage(package_name)
			if err != nil{
				if (err != requester.ErrPageNotFound){
					return nil, err			
				}
				isFound = false
			}
			
			var formattedResponse RegistryPackage
			json.Unmarshal(response, &formattedResponse)
			
			// Initializing Empty Values
			// If not initalized it'll cause `assign to entry in nil map` error
			formattedResponse.Dependencies = make(map[string]RegistryPackageDependencyInfo)
			formattedResponse.NPMUsers = []string{}

			if latestVersion, ok := formattedResponse.DistTags["latest"]; ok{
				// Loop through formattedResponse.Versions keys and find the version that is equal to the LatestVersion
				for version, pkgValue := range formattedResponse.Versions{
					if version == latestVersion {
						for dependency, version := range pkgValue.Dependencies{
							if _, ok := formattedResponse.Dependencies[dependency]; !ok{
								_, err := n.LookupPackage(dependency)
								formattedResponse.Dependencies[dependency] = RegistryPackageDependencyInfo{
									Version: version,
									Exists: (err != requester.ErrPageNotFound),
								}
							}
						}

						for dependency, version := range pkgValue.DevDependencies{
							if _, ok := formattedResponse.Dependencies[dependency]; !ok{
								_, err := n.LookupPackage(dependency)
								formattedResponse.Dependencies[dependency] = RegistryPackageDependencyInfo{
									Version: version,
									Exists: (err != requester.ErrPageNotFound),
								}
							}
						}
					}

					if !Contains(formattedResponse.NPMUsers, pkgValue.NPMUser.Email){
						formattedResponse.NPMUsers = append(formattedResponse.NPMUsers, pkgValue.NPMUser.Email)
					}
				}
			}

			formattedResponse.Versions = nil

			return map[string]interface{}{
				"Content": formattedResponse,
				"Vulnerable": !isFound,
			}, nil
		case source.URLSource:
			
			// This adapter only responds to "npm" Registry meaning that any URL
			// Sent to this adapter has to be validated through the SourceIdentifier
			// Enforcing that it's a PackageJSON based on JSON Response and structure
			// Therefore we'll be expected PackageJSON, and i'll be creating another
			// Adapter that will use core utils to Crawl and Parse "Unknown" Registry Types
			var URLPackageContent NPMPackageJSON
			// var srcIdentifierFileType string
			var nfPkgs []map[string]string
			var fPkgs []map[string]string
			
			// We need to access an "undefined" attribute that exists in the Metadata
			o := reflect.ValueOf(src.Metadata)
			// map[string]interface{} - map[string]string
			if o.Kind() == reflect.Map{
				for _, key := range o.MapKeys(){
					switch strings.ToLower(key.String()){
						case "content":
							// .String()
							// Unlike the other getters, it does not panic if v's Kind is not String.
							// Instead, it returns a string of the form "<T value>" where T is v's type
							// content := o.MapIndex(key).String()
							// .Interface()
							// It panics if the Value was obtained by accessing unexported struct fields
							content := o.MapIndex(key).Interface()
							if contentStr, ok := content.(string); ok {
								err := json.Unmarshal([]byte(contentStr), &URLPackageContent)
								if err != nil {
									return nil, err
								}
							}
						// case "filetype":
							// srcIdentifierFileType = o.MapIndex(key).String()
						default:
							// fmt.Println(key.String())
					}
				}
			}

			// Extending Dependencies with DevDependencies
			for devPkg, devPkgVersion := range URLPackageContent.DevDependencies{
				URLPackageContent.Dependencies[devPkg] = devPkgVersion
			}

			for pkgName, pkgVersion := range URLPackageContent.Dependencies{
				_, err := n.LookupPackage(pkgName)
				if err != nil && err == requester.ErrPageNotFound{
					// What if a package is found but a different error Occured?
					// idk i need to create Tests
					nfPkgs = append(nfPkgs, map[string]string{pkgName: pkgVersion})
				}else{
					fPkgs = append(fPkgs, map[string]string{pkgName: pkgVersion})
				}
			}

			return map[string][]map[string]string{
				"Not_Found": nfPkgs,
				"Found": fPkgs,
			}, nil
		default:
			return nil, nil
	}
}

func (n *RegistryModule) CreatePackage(args map[string]interface{}) error {
	package_name := args["packageName"].(string)

	fmt.Printf("Creating NPM package for %s\n", package_name)

	payloadFiles := n.GithubAPI.GetPayloadFiles(payloads.DemoPayload)
	for _, file := range payloadFiles{
		fmt.Println(file.Path)
	}

	n.channel<-package_name
	return nil
}
