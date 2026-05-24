package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/v2fly/geoip/lib"
)

var (
	// configFile is the path to the JSON configuration file
	configFile = flag.String("config", "config.json", "Path to the configuration file")
	// outputDir overrides the output directory specified in config
	outputDir = flag.String("output", "", "Override output directory from config")
	// listInputs prints all available input formats and exits
	listInputs = flag.Bool("list-inputs", false, "List all available input formats")
	// listOutputs prints all available output formats and exits
	listOutputs = flag.Bool("list-outputs", false, "List all available output formats")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nA tool to generate GeoIP data files in various formats.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExample:\n")
		fmt.Fprintf(os.Stderr, "  %s -config config.json\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -config config.json -output ./output\n", os.Args[0])
	}

	flag.Parse()

	if *listInputs {
		lib.PrintAvailableInputs()
		os.Exit(0)
	}

	if *listOutputs {
		lib.PrintAvailableOutputs()
		os.Exit(0)
	}

	// Load and validate the configuration file
	instance, err := lib.NewInstance()
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Failed to create instance: %v\n", err)
		os.Exit(1)
	}

	if err := instance.LoadConfigFile(*configFile); err != nil {
		fmt.Fprintf(os.Stderr, "❌ Failed to load config file %q: %v\n", *configFile, err)
		os.Exit(1)
	}

	// Override output directory if specified via flag
	if *outputDir != "" {
		instance.SetOutputDir(*outputDir)
	}

	fmt.Printf("🌍 GeoIP data generator starting...\n")
	fmt.Printf("📄 Config: %s\n", *configFile)

	// Run the full pipeline: input -> process -> output
	if err := instance.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "❌ Generation failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ GeoIP data files generated successfully.\n")
}
