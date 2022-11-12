# gotrtc
Go package for ctrtc(tencent trtc)

## Use

### install ctrtc

```
git clone https://github.com/lynnplus/ctrtc.git
cmake -S ${ctrtc_dir} -B ${ctrtc_build_dir}
cmake install
```

### use gotrtc
`go get github.com/lynnplus/gotrtc`

example:

```
func main(){
    trtc := gotrtc.GetShareInstance()
    fmt.Println("sdk version:", trtc.GetSDKVersion())
    gotrtc.DestroyShareInstance()
}

```

run:
```
set CGO_CPPFLAGS=-I${ctrtc_install_dir}\include
set CGO_LDFLAGS=-L${ctrtc_install_dir}\lib
```

Copy the files in the `${ctrtc_install_dir}\bin` folder to your bin(run) directory

`go run main.go`
