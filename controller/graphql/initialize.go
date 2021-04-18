package graphql

type GraphqlController struct{}

func Initialize(defaultPort string) {
	var gc GraphqlController
	gc.register(defaultPort)
}
