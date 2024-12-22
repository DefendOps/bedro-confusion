package npm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/defendops/bedro-confuser/pkg/registry"
	"github.com/defendops/bedro-confuser/pkg/utils/requester"
	"github.com/defendops/bedro-confuser/pkg/utils/source"
)

type RegistryModule struct {}

func init() {
	npmModule := RegistryModule{}
	registry.RegisterRegistryModule(&npmModule)
}

func (n *RegistryModule) Run(src source.Source, ctx *context.Context) error {
	adaptedSource, err := n.SourceAdapter(src)
	
	switch err {
		case requester.ErrPageNotFound:
			// Package not found, maybe vulnerable.
			switch src.Type{
				case source.PackageNameSource:
					fmt.Println("Package Not Found, Maybe vulnerable")
				case source.URLSource:
					fmt.Println("URL is returning 404?")
				default:
					fmt.Println("WTF?")
			}
		case nil:
			// Package Found, not vulnerable
			fmt.Println(adaptedSource)
			return nil
		default:
			return errors.New("i don't really know what happened")
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

func (n *RegistryModule) SourceAdapter(src source.Source) (interface{}, error) {
	switch src.Type{
		case source.PackageNameSource:
			package_name := src.RawValue
			package_name = strings.ToLower(package_name)

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

			response, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading URL response: %v", err)
			}
			
			var formattedResponse RegistryPackage
			err = json.Unmarshal(response, &formattedResponse)
			if err != nil{
				return nil, err
			}

			return map[string]RegistryPackage{"content": formattedResponse}, nil
		case source.URLSource:
			fmt.Println(src)
			return nil, nil
		default:
			return nil, nil
	}
}

func (n *RegistryModule) Scan(adaptedSource interface{}) error {
	fmt.Println("Scanning")
	return nil
}

func (n *RegistryModule) CreatePackage() error {
	fmt.Println("Creating NPM package...")
	return nil
}
