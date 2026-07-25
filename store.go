package main

import (
	"sync"
	"time"
)
 
var store = make(map[string]string)

var expire = make(map[string]time.Time)

var mu sync.RWMutex