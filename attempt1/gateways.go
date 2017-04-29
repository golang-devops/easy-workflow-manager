package attempt1

type Gateway struct{}
type ExclusiveGateway struct{ Gateway }
type EventBasedGateway struct{ Gateway }
type ParallelGateway struct{ Gateway }
type InclusiveGateway struct{ Gateway }
type ExclusiveEventBasedGateway struct{ Gateway }
type ComplexGateway struct{ Gateway }
type ParallelEventBasedGateway struct{ Gateway }
