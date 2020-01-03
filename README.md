[![GitHub](https://img.shields.io/github/license/gopowersupply/intelcpu)](https://intelcpu/blob/master/LICENSE)
[![Go version](https://img.shields.io/github/go-mod/go-version/gopowersupply/intelcpu)](https://blog.golang.org/go1.13)
[![Go Report Card](https://goreportcard.com/badge/gopowersupply/intelcpu)](http://goreportcard.com/report/gopowersupply/intelcpu)
[![code-coverage](http://gocover.io/_badge/github.com/gopowersupply/intelcpu)](https://gocover.io/github.com/gopowersupply/intelcpu)
[![GoDoc](https://godoc.org/github.com/gopowersupply/intelcpu?status.svg)](https://godoc.org/github.com/gopowersupply/intelcpu)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/gopowersupply/intelcpu)](https://github.com/gopowersupply/intelcpu/releases)
[![GitHub last commit](https://img.shields.io/github/last-commit/gopowersupply/intelcpu)](https://intelcpu/commits/master)
[![GitHub issues](https://img.shields.io/github/issues/gopowersupply/intelcpu)](https://intelcpu/issues)

The package to allow interaction with intel_pstate driver to drive CPU frequency and politics.  

**intel_pstate** driver must be installed and activated. Usually, it __already included in Linux kernel 4.13 for Sandy Bridge and later__ CPUs.

Get it from github:
```bash
go get -u https://github.com/gopowersupply/intelcpu
```

Documentation can be [found here](https://godoc.org/github.com/gopowersupply/intelcpu)

# Requirements

- `intel_pstate` CPU Performance Scaling Driver

# Examples

Simple example to change TurboBoost status:
```go
    cpu := intelcpu.New()
    turbo, _ := cpu.GetTurbo()
    if turbo {
    	cpu.SetTurbo(false)
    } else {
    	cpu.SetTurbo(true)
    }
```

In real projects strongly recommended to check for driver and its status:
```go
    cpu := intelcpu.New()
    
    if err := cpu.CheckDriver(); err != nil {
    	// [...] Some troubles or driver not installed
    }
    
    status, _ := cpu.GetStatus()
    switch status {
    case intelcpu.PStateStatusActive:
    	// [...] All is ok
    case intelcpu.PStateStatusPassive:
    	// [...] Something wrong, working partially
    case intelcpu.PStateStatusOff:
    	// [...] Driver disabled, nothing to work
    }
```

You can enable or disable some cores. Except first, of course:
```go
    cpu := intelcpu.New()
        
    cores, _ := cpu.GetCores()
    
    for _, core := range cores {
    	// First core will return false and its status always will be online
    	isOfflineAvailable, _ := core.IsOfflineAvailable()
    	    	
    	isOnline, _ := core.IsOnline()
    	fmt.Printf("Core %d is online: %v", isOnline)
    	
    	// If core can be offline then do it
    	if isOfflineAvailable {
    		core.SetOnline(false)
    	}    	
    }
```

You can change CPU frequency limitation also:
```go
    cpu := intelcpu.New()    
    cpu.SetMaxPerf(0.5) // 50% of max
```

Core performance and governor politics also can be changed:
```go
    cpu := intelcpu.New()
    cores, _ := cpu.GetCores()
    
    for _, core := range cores {
    	core.SetGovernor(intelcpu.CPUGovernorPerformance)
    	core.SetPreference(intelcpu.CPUPreferencePerformance)
    }
```

Short way:
```go
    cpu := intelcpu.New()
    cores, _ := cpu.GetCores()
    cores.SetGovernor(intelcpu.CPUGovernorPerformance)
    cores.SetPreference(intelcpu.CPUPreferencePerformance)
```

## Errors handling

This package has an own error type `CPUError`  
You can pass the package errors through your functions then detect it via `errors.As`:
```go
    func ExecUnexpected() error {
    	// [...] Here your other returns with own errors
        cpu := intelcpu.New()
        _, err := cpu.GetCore(20000)
        if err != nil {
        	return err
        }
        // [...] Here your other returns with own errors
    }

    func main() {
    	err := ExecUnexpected()    	
    	if intelcpu.IsCPUError(err) {
    		// [...] to do anything
    	} else {
    		// [...] to do something other    		
    	}
    }
```
And you can use `errors.As(err, &intelcpu.CPUError{})` as alternative.