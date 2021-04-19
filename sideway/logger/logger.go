package main

import (
	"go.uber.org/zap"
)

func main() {
	l := zap.NewExample()
	defer l.Sync()
	l = l.With(zap.Namespace("hometicx"), zap.String("I'm", "Gopher"))
	l.Info("Hello")
}
