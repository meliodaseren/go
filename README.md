## GO GO POWER (RANGERS)

> 跟風學一下，打發時間啦。

**OS**: macOS 10.13 / Windows 7 64 bits

**Version**: Go 1.12.4

**IDE**: VS Code 1.33 / ~~GoLand 2017.3.2~~

### Tutorial

https://tour.golang.org/

https://github.com/astaxie/build-web-application-with-golang/blob/master/zh-tw/preface.md

### How to set the environment ?

> 因為今天比較趕，速速紀錄一下今天在 Windows 7 的環境設置。

#### 1. Set the environment variable

`GOROOT` : Default may be `C:\Go\`

`GOPATH` : Default may be `C:\gotool\`

#### 2. Install the required packages

```sh
go get -u -v github.com/nsf/gocode  
go get -u -v github.com/rogpeppe/godef  
go get -u -v github.com/golang/lint/golint  
go get -u -v github.com/lukehoban/go-find-references  
go get -u -v sourcegraph.com/sqs/goreturns  
go get -u -v golang.org/x/tools/cmd/gorename  
go get -u -v github.com/derekparker/delve/cmd/dlv 
# go get -u -v github.com/zmb3/gogetdoc
# go get -u -v github.com/ramya-rao-a/go-outline
# go get -u -v github.com/tpng/gopkgs
# go get -u -v github.com/acroca/go-symbols
# go get -u -v golang.org/x/tools/cmd/guru
# go get -u -v github.com/cweill/gotests/...
# go get -u -v golang.org/x/tools/cmd/godoc
# go get -u -v github.com/fatih/gomodifytags
# go get -u -v github.com/josharian/impl 
```

#### 3. Install the VS Code & Extension

Type `Ctrl + Shift + P` and find the Go extension, then install it

Go Extension: https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go

#### 4. Modify the `launch.json` for running and debugging

```json
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

#### 5. Type the `F5` for running and debugging
