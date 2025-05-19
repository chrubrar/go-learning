package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// Config demonstrates the use of sync.Once for lazy initialization.
type Config struct {
	settings map[string]string
	once     sync.Once
}

// NewConfig creates a new Config instance with an initialized settings map.
func NewConfig() *Config {
	return &Config{
		settings: make(map[string]string),
	}
}

// LoadSettings loads configuration settings exactly once, even when called
// multiple times or from multiple goroutines.
func (c *Config) LoadSettings() {
	c.once.Do(func() {
		fmt.Println("Loading settings... (this should print only once)")
		// Simulate some heavy initialization
		time.Sleep(time.Second)
		
		c.settings["database"] = "mysql"
		c.settings["host"] = "localhost"
		c.settings["port"] = "3306"
	})
	
	fmt.Println("Settings ready to use")
}

// GetSetting returns the value for a given setting key.
func (c *Config) GetSetting(key string) string {
	return c.settings[key]
}

// DemonstrateSyncOnce shows how sync.Once ensures initialization happens exactly once
// even when called from multiple goroutines.
func DemonstrateSyncOnce() {
	config := NewConfig()
	var wg sync.WaitGroup
	
	// Launch 5 goroutines that all try to load settings
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d attempting to load settings...\n", id)
			config.LoadSettings()
			// Read a setting to prove it's loaded
			fmt.Printf("Goroutine %d sees database=%s\n", id, config.GetSetting("database"))
		}(i)
	}
	
	wg.Wait()
}
