package driver

import (
	"sync"
	logger "github.com/AdityaSubrahmanyaBhat/dashDB/models/Logger"
)

type Driver struct {
	Dir string
	Mutex   sync.Mutex
	Mutexes map[string]*sync.Mutex
	Log     logger.Logger
}