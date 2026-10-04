package inmemory_test

import (
	"testing"

	"go.trueardian.com/chronica/inmemory"
	"go.trueardian.com/chronica/storeconformance"
)

func TestConformance(t *testing.T) {
	storeconformance.RunTest(t, inmemory.NewStore)
	storeconformance.RunIdempotentTest(t, inmemory.NewIdempotentStore)
}
