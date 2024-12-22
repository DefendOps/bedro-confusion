package registry

import (
	"fmt"
	"sync"
)

var (
	RegistryModules = make(map[Registry][]RegistryModule)
	moduleIDs       = make(map[RegistryModuleID]bool)
	mutex           = sync.Mutex{} // For thread-safe operations
)

func RegisterRegistryModule(r RegistryModule) {
	id := CreateRegistryModuleID(r)

	mutex.Lock()
	defer mutex.Unlock()

	if moduleIDs[id] {
		panic(fmt.Sprintf("module: Register called twice for module %+v\n", id))
	}

	moduleIDs[id] = true

	var moduleList []RegistryModule
	if list, exists := RegistryModules[r.Registry()]; exists {
		moduleList = list
	} else {
		moduleList = make([]RegistryModule, 0)
	}

	RegistryModules[r.Registry()] = append(moduleList, r)
}

func CreateRegistryModuleID(r RegistryModule) RegistryModuleID {
	return RegistryModuleID{
		name:     r.Name(),
		registry: r.Registry(),
	}
}

func (rid RegistryModuleID) String() string {
	return fmt.Sprintf("%v", rid.name)
}
