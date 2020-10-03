package evolver


type Evolver interface{
	Evolve(string, *string) error
}

func NewEvolver(edition string, evolverType string) Evolver{
	switch (edition){
		case "lv1":
			switch (evolverType){
			case "PHYSICALNODE":
				evolver :=new(PhysicalNode)
				return evolver
			default:
				return nil
			}
		default:
			return nil
	}
}