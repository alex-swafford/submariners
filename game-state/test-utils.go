package game_state

import (
	"fmt"
	"testing"
)

func assertEquals(value1 any, value2 any, t *testing.T) {
	if value1 != value2 {
		t.Errorf("Expected " + fmt.Sprint(value1) + ". Got " + fmt.Sprint(value2) + ".")
		t.Fail()
	}
}

func assertNil(value any, t *testing.T) {
	if value != nil {
		t.Errorf("Expected " + fmt.Sprint(value) + " to be nil.")
		t.Fail()
	}
}

func assertNotNil(value any, propertyName string, t *testing.T) {
	if value == nil {
		t.Errorf("Expected " + propertyName + " to not be nil, but it was.")
		t.Fail()
	}
}
