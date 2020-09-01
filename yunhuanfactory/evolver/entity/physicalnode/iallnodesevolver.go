package physicalnode


type IAllNodesEvolver interface{
	Evolver(string, string) string
}