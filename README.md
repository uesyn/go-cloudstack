gokcps
=============
This package is based on [xanzy/go-cloudstack](https://github.com/xanzy/go-cloudstack).

## Description
A KDDI Cloud Platform Service API client enabling Go programs to interact with KCPS in a simple and uniform way.

## Example

```go
// Create a new API client
cli := gokcps.NewAsyncClient("api_endpoint", "api_key", "secret_key" , false) 

// Create a new parameter struct.  
// If you don't specify a specific network, 'PublicFrontSegment' is used as network that is used for virtual machine.
p := cli.VirtualMachine.NewDeployValueVirtualMachineParams(
	"service-offering-id", 
	"template-id", 
	"zone-id",
	"vmname",
)

// If you want to change or add some parameters,
// You can set with following parameter's method.
// Set any other options required by your setup in the same way.
p.SetName("vmname2")

// Create the new instance
r, err := cli.VirtualMachine.DeployValueVirtualMachine(p)
if err != nil {
	log.Fatalf("Error creating the new instance %s: %s", name, err)
}

fmt.Printf("UUID or the newly created machine: %s", r.ID)
```

## ToDo
- [ ] godoc
- [ ] implementation of some convinient method. 

## Getting Help
* If you have an issue: report it on the [issue tracker](https://github.com/uesyn/gokcps/issues)

## Author
[uesyn](https://github.com/uesyn)

## License
[Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0)

