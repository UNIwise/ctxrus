# ctxrus

A very small helper package to store a logrus entry in a context. The reason for this project is that we end up passing a context with around anyway and it would be nice having context based logging fields when logging information.

### Usage

More usage examples can be found in the examples folder.
```go
func main() {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})
	ctxrus.SetLogger(l)

	ctx := ctxrus.AddFields(context.Background(), logrus.Fields{
		"method": "GET",
	})

	myFunc(ctx)
}

func myFunc(ctx context.Context) {
	ctxrus.GetLogger(ctx).Info("hello") // Log: {"level":"info","method":"GET","msg":"hello","time":"2022-06-30T13:02:14+02:00"}
}
```