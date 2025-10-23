package testdemo

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setUp() {
	// do setup
	fmt.Println("setup completed")
}

func tearDown() {
	// do clean
	fmt.Println("teardown completed")
}

func testSetup(tb testing.TB) func(t testing.TB) {
	fmt.Println("test setup")

	return func(tb testing.TB) {
		fmt.Println("test teardown")
	}
}

func TestDoSomething(t *testing.T) {
	type data struct {
		name string
		input int
		except int
	}

	testData := []data{
		{
			name: "test1",
			input: 3,
			except: 1,
		},
		{
			name: "test2",
			input: 8,
			except: 2,
		},
	}

	for _, tdata := range testData {
		t.Run(tdata.name, func(t *testing.T) {
			teardown := testSetup(t)
			defer teardown(t)

			if tdata.input != tdata.except {
				t.Errorf("test assert demo failed")
			}
		})
	}
}
