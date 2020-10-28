package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Permission is a type alias for int and works with Iota to create an enum of all possible permissions.
type Permission int

// These are the possible permissions a file can have.
const (
	R Permission = iota
	W
	X
	RW
	RX
	WX
	RWX
)

func (p Permission) String() string {
	return [...]string{"R", "W", "X", "RW", "RX", "WX", "RWX"}[p]
}

// File represents a file (unique to a user) that includes the name and that user's permissions.
type File struct {
	Name string
	Perm Permission
}

// NamePerm represents a user and permission to a file (unique to a file).
type NamePerm struct {
	User string
	Perm Permission
}

func main() {
	acl := map[string][]File{}
	aclByFile := map[string][]NamePerm{}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100000; i++ {
		perm := Permission(rand.Intn(6))
		if i%1000 == 0 {
			acl["user"+strconv.Itoa(i)] = []File{{"world", perm}}
			aclByFile["world"] = append(aclByFile["world"], NamePerm{"user" + strconv.Itoa(i), perm})
			continue
		}
		acl["user"+strconv.Itoa(i)] = []File{{"hello", perm}}
		aclByFile["hello"] = append(aclByFile["hello"], NamePerm{"user" + strconv.Itoa(i), perm})
	}

	fmt.Printf("user75432 has access to file %v with permissions %v\n", acl["user75432"][0].Name, acl["user1005"][0].Perm)
	fmt.Printf("All users that have access to file 'world': %v\n", aclByFile["world"])
}
