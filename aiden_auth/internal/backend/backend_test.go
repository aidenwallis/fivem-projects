package backend_test

import (
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/backend"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/config"
	"github.com/aidenwallis/fivem-projects/aiden_auth/internal/fakes"
)

const (
	// fakeToken is not a real token, it's used in test inputs
	fakeToken = "abc123"
	// hashedFakeToken is a sha256 hash of fakeToken, again, not a real token that we would generate
	hashedFakeToken = "bKE9UspwyIPg8LsQHkJaiehiTeUdstI5JZOvaoQRgJA"

	// how long our tests pretend that tokens last
	lifetimeSeconds = 300
)

func newTestEnvironment() (backend.Backend, *fakes.FakeDB, *fakes.FakeLogger) {
	db := &fakes.FakeDB{}
	log := &fakes.FakeLogger{}
	return backend.NewBackend(db, log, &config.SessionsConfig{
		LifetimeSeconds: lifetimeSeconds,
	}), db, log
}
