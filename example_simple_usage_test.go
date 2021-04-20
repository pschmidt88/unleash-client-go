package unleash_test

import (
	"fmt"
	"time"

	"github.com/pschmidt88/unleash-client-go/v3"
)

const PropertyName = "eid.enabled"

// ExampleSimpleUsage demonstrates the simplest way to use the unleash client.
func Example_simpleUsage() {
	unleash.Initialize(
		unleash.WithListener(&unleash.DebugListener{}),
		unleash.WithAppName("my-application"),
		unleash.WithUrl("http://unleash.herokuapp.com/api/"),
	)

	timer := time.NewTimer(1 * time.Second)

	for {
		<-timer.C
		fmt.Printf("'%s' enabled? %v\n", PropertyName, unleash.IsEnabled(PropertyName))
		timer.Reset(1 * time.Second)
	}

}
