# ctxrus

A very small helper package to store a logrus entry in a context. The reason for this project is that we end up passing a context with around anyway and it would be nice having context based logging fields when logging information.

### Usage

More usage examples can be found in the examples folder.
```go
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
```