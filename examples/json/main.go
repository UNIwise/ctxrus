package main

import (
	"context"

	"github.com/UNIwise/ctxrus"
	"github.com/sirupsen/logrus"
)

func main() {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})
	ctxrus.SetLogger(l)

	myFunc()
}

func myFunc() {
	ctx := context.Background()
	ctxrus.GetLogger(ctx).Info("hello")
}
