package main

import (
	"context"

	"github.com/UNIwise/ctxrus"
	"github.com/sirupsen/logrus"
)

func main() {
	l := logrus.New()
	l.SetFormatter(&logrus.TextFormatter{})
	ctxrus.SetLogger(l)

	myFunc()
}

func myFunc() {
	ctx := ctxrus.AddFields(
		context.Background(),
		logrus.Fields{
			"trace_id": "0fbc8d98-9bac-49f1-b86b-cd9117a60bd0",
		},
	)
	myNestedFunc(ctx)
}

func myNestedFunc(ctx context.Context) {
	ctxrus.GetLogger(ctx).Info("hello")
	myNestedFunc2(ctxrus.AddFields(ctx, logrus.Fields{
		"method": "GET",
	}))
}

func myNestedFunc2(ctx context.Context) {
	ctxrus.GetLogger(ctx).Info("hello")
}
