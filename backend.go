package backend

import (
	"github.com/sirupsen/logrus"
)

// Config struct contains all the essential information the application
// should have available on all scopes
type Config struct {
	Development bool

	Logger         *logrus.Logger
	ExecutablePath string

	// Public Address
	PublicHost string
	Host       string
	Port       int
	Storage    string
}
