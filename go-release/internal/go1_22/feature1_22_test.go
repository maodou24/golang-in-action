package go1_22

import (
	"fmt"
	"io"
	"math/rand/v2"
	"net/http/httptest"
	"testing"
)

func TestIterateSliceInClose(t *testing.T) {
	IterateSliceInClose()
}

func TestMathRandV2IntNSourceChaCha8(t *testing.T) {
	r := rand.New(rand.NewChaCha8([32]byte([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ123456"))))
	for i := 0; i < 10; i++ {
		fmt.Println(r.IntN(100))
	}
}

func TestMathRandV2IntNSourcePCG(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 2))
	for i := 0; i < 10; i++ {
		fmt.Println(r.IntN(100))
	}
}

func TestMathRandV2N(t *testing.T) {
	fmt.Println(rand.N(100))
}

func TestIsGoVersionOlder(t *testing.T) {
	fmt.Println(IsGoVersionOlder("go1.21", "go1.22"))
	fmt.Println(IsGoVersionOlder("go1.22", "go1.21"))
}

func TestHttp(t *testing.T) {
	srv := httptest.NewServer(RoutingPatternsServeMux())

	cli := srv.Client()

	resp, err := cli.Get(srv.URL + "/items/1")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if string(body) != "1" {
		t.Errorf("Expected id is 1, got %s", string(body))
	}

	resp, err = cli.Get(srv.URL + "/files/file1/file2/file3")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	if string(body) != "file1/file2/file3" {
		t.Errorf("Expected id is file1/file2/file3, got %s", string(body))
	}
}
