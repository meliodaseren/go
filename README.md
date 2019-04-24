## Go Go Power

> 跟風學一下，打發時間啦。

OS: macOS 10.13 / Windows 7

Version: Go 1.12.4

IDE: VS Code 1.33

     GoLand 2017.3.2 (Dropped)

### Tutorial

https://tour.golang.org/

### How to set the environment ?

1. Setting the `GOROOT` & `GOPATH`

`GOROOT` : Default may be `C:\Go\`

`GOPATH` : Default may be `C:\gotool\`

2. Install the required packages

```sh
go get -u -v github.com/nsf/gocode  
go get -u -v github.com/rogpeppe/godef  
go get -u -v github.com/golang/lint/golint  
go get -u -v github.com/lukehoban/go-find-references  
go get -u -v sourcegraph.com/sqs/goreturns  
go get -u -v golang.org/x/tools/cmd/gorename  
go get -u -v github.com/derekparker/delve/cmd/dlv  
```

3. Install the VS Code & Extension

Type `Ctrl + Shift + P` and find the `Go` extension, then install it.

Go Extension: https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go

4. Modify the `launch.json` for running and debugging

```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            //"mode": "deubg",
            //"program": "${fileDirname}",
            "program": "${workspaceRoot}",
            "env": {},
            "args": []
        }
    ]
}
```

5. Type the `F5` for running and debugging
