package uuid_test

import (
	"fmt"

	"github.com/pschlump/uuid"
)

func ExampleNewV4() {
	u4, err := uuid.NewV4()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%s\n", u4)
}

func ExampleNewV5() {
	u5, err := uuid.NewV5(uuid.NamespaceURL, []byte("nu7hat.ch"))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("%s\n", u5)
}

func ExampleParseHex() {
	u, err := uuid.ParseHex("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("%s\n", u)
}
