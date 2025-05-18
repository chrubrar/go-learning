package main

import (
	"fmt"
	"sync"
	"time"
)

// Config represents a configuration that should only be loaded once
type Config struct {
	settings map[string]string
	once     sync.Once
}

func NewConfig() *Config {
	return &Config{
		settings: make(map[string]string),
	}
}

func (c *Config) LoadSettings() {
	// sync.Once ensures this code runs exactly once, even if called from multiple goroutines
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

func main() {
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
			fmt.Printf("Goroutine %d sees database=%s\n", id, config.settings["database"])
		}(i)
	}
	
	wg.Wait()
	fmt.Println("\nFinal settings:", config.settings)
}
