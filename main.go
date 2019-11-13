package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
	"runtime"

	logger "github.com/bjr3ady/go-logger"
	models "git.r3ady.com/golang/school-board/models/orm"
	"git.r3ady.com/golang/school-board/pkg/setting"
	
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	logger.DefaultPrefix = setting.RunMode
}