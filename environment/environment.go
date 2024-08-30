package environment

import (
	"os"
	"runtime"
)

const (
	EnvWindows = "windows"
	EnvLinux   = "linux"
)

type Env struct {
	CurDir   string
	Platform string
	Separate string
}

func Load() *Env {

	var e Env
	// 加载系统变量
	curDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	e.CurDir = curDir

	if runtime.GOOS == EnvWindows {
		e.Separate = "\\"
	} else {
		e.Separate = "/"
	}

	return &e
}
