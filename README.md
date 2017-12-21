gokcps
=============
This package is based on [xanzy/go-cloudstack](https://github.com/xanzy/go-cloudstack).

## Description
A KDDI Cloud Platform Service API client enabling Go programs to interact with KCPS in a simple and uniform way.

## Example

```go
// API Clientを生成する
cli := gokcps.NewAsyncClient("api_endpoint", "api_key", "secret_key" , false) 

// APIを叩くためのパラメータを生成する.  
// gokcpsでKCPSのAPIを叩く場合は対称のAPIのパラメータオブジェクトを生成し、そのパラメータオブジェクトを実際にAPIを叩くメソッドの引数とする.
// パラメータオブジェクト生成するメソッドは全てNew***Paramsの命名規則を持つ.
// パラメータオブジェクト生成時にはAPIを叩くために必須のパラメータを引数として渡す必要がある(一部例外あり)
p := cli.VirtualMachine.NewDeployValueVirtualMachineParams(
	"service-offering-id", 
	"template-id", 
	"zone-id",
	"vmname",
)

// パラメータオブジェクトにはAPIで利用可能なその他のオプションをSet***メソッドで追加可能.
p.SetName("vmname2")

// 上記で生成したパラメータを引数にAPIを実行.
// 全てのAPI実行メソッドには一つ目の返り値にResponseオブジェクトと、二つ目の返り値に実行エラーが帰ってくる.(成功時にはnil)
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

