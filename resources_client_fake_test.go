// nolint:errcheck
package modzy

import (
	"context"
	"testing"
)

func TestResourcesClientFake(t *testing.T) {
	expectedCtx := context.WithValue(context.TODO(), testContextKey("a"), "b")

	calls := 0
	fake := &ResourcesClientFake{
		GetProcessingModelsFunc: func(ctx context.Context) (*GetProcessingModelsOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			return nil, nil
		},
	}

	fake.GetProcessingModels(expectedCtx)

	if calls != 1 {
		t.Errorf("Did not call all of the funcs: %d", calls)
	}
}
