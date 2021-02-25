package welcome

import "testing"

func TestHelloEmptyName(t *testing.T) {
	val, err := sayHello("")

	if len(val) != 0 || err == nil {
		t.Fatalf("Expected val = \"\" and  err !=nil, but was %s,%v", val, err)
	}
}

func TestHello(t *testing.T) {
	val, err := sayHello("Andreea")

	if val != "Hello Andreea" || err != nil {
		t.Fatalf("Expected val = Hello Andreea and err = nil, but was %s,%v", val, err)
	}
}
