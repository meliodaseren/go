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

#### 4. Modify the setting file

*launch.json*:

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            //"name": "Launch",
            "name": "dlv-DEBUG",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            //"mode": "debug",
            //"program": "${fileDirname}",
            "program": "${workspaceRoot}",
            "windows": {
                "env": {
                    "GOPATH": "${env:GOPATH};${workspaceRoot}"
                }
            },
            "osx": {
                "env": {
                    "GOPATH": "${env:GOPATH}:${workspaceRoot}"
                }
            },
            "linux": {
                "env": {
                    "GOPATH": "${env:GOPATH}:${workspaceRoot}"
                }
            },
            "args": []
        }
    ]
}
```

*settings.json*:

```json
{
    /* Golang */
    "files.autoSave": "onFocusChange",
    "go.buildOnSave": "true",
    "go.lintOnSave": "true",
    "go.vetOnSave": "true",
    "go.buildFlags": [],
    "go.lintFlags": [],
    "go.vetFlags": [],
    "go.useCodeSnippetsOnFunctionSuggest": false,
    "go.formatTool": "goreturns",
    "go.goroot": "/usr/local/opt/go/libexec",
    "go.gopath": "${your_project_folder}",
    // User defined GOPATH setting
    //"terminal.integrated.shell.windows": "",
    //"go.gopath": "${workspaceRoot}",
    //"go.inferGopath": true
}
```

*tasks.json*:

```json
{
    // See https://go.microsoft.com/fwlink/?LinkId=733558
    // for the documentation about the tasks.json format
    "version": "2.0.0",
    "windows": {
        "options": {
            "env": {
                // ${workspaceRoot} 當前打開的文件夾的絕對路徑
                "GOPATH": "${env:GOPATH};${workspaceRoot}"
            }
        }
    },
    "osx": {
        "options": {
            "env": {
                "GOPATH": "${env:GOPATH}:${workspaceRoot}"
            }
        }
    },
    "linux": {
        "options": {
            "env": {
                "GOPATH": "${env:GOPATH}:${workspaceRoot}"
            }
        }
    },
    "tasks": [
        {
            "label": "go run",
            "command": "go",
            "type": "shell",
            "group": "none",
            "args": [
                "run",
                //"main.go"
                "${file}"
                //${fileBasename}
            ],
            "promptOnClose": false,
            "presentation": {
                "reveal": "always"
            },
            "problemMatcher": [
                "$go"
            ]
        },
        {
            "label": "go build",
            "command": "go",
            "type": "shell",
            "promptOnClose": false,
            "group": {
                "kind": "build",
                "isDefault": true
            },
            "args": [
                "build",
                "-i",
                "-v",
                //"${workspaceRootFolderName}",
                "${file}",
                //"|",
                "&&",
                //"./${workspaceRootFolderName}"
                "./${fileBasenameNoExtension}"
            ],
            "presentation": {
                "reveal": "always"
            },
            "problemMatcher": [
                "$go"
            ]
        },
        {
            "label": "go build & exec",
            "command": "go",
            "type": "shell",
            "group": "none",
            "windows": {
                "args": [
                    "build",
                    "-i",
                    "-v",
                    //"${workspaceRootFolderName}",
                    "${file}",
                    "&",
                    //"${workspaceRootFolderName}.exe"
                    "${fileBasenameNoExtension}.exe"
                ]
            },
            "linux": {
                "args": [
                    "build",
                    "-i",
                    "-v",
                    //"${workspaceRootFolderName}",
                    "${file}",
                    "|",
                    //"./${workspaceRootFolderName}"
                    "./${fileBasenameNoExtension}.exe"
                ]
            },
            "osx": {
                "args": [
                    "build",
                    "-i",
                    "-v",
                    //"${workspaceRootFolderName}",
                    "${file}",
                    //"|",
                    "&&",
                    //"./${workspaceRootFolderName}"
                    "./${fileBasenameNoExtension}.exe"
                ]
            },
            "promptOnClose": false,
            "presentation": {
                "reveal": "always"
            },
            "problemMatcher": [
                "$go"
            ]
        }
    ]
}
```

#### 5. Type the `F5` for running and debugging
