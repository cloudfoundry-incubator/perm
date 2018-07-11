package rpc_test

import (
	. "code.cloudfoundry.org/perm/api/internal/rpc"

	"code.cloudfoundry.org/perm/api/internal/repos"
	. "code.cloudfoundry.org/perm/api/internal/repos/reposbehaviors"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("InMemoryStore", func() {
	var (
		store *InMemoryStore
	)
	BeforeEach(func() {
		store = NewInMemoryStore()
	})

	BehavesLikeARoleRepo(func() repos.RoleRepo { return store })
	BehavesLikeAPermissionRepo(
		func() repos.PermissionRepo { return store },
		func() repos.RoleRepo { return store },
	)
})
