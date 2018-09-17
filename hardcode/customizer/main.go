package customizer

import (
	"fmt"
	"os"
)

// tag::options[]
type Options struct {
	AccessKeyID     string
	SecretAccessKey string
	BucketName      string
	BucketPath      string
}

// end::options[]

// tag::vars[]
var opt Options
var globalConfig string

// end::vars[]

// tag::mainopt[]
func Main(o Options) {
	opt = o
	// end::mainopt[]

	// tag::mainfetch[]
	err := fetch(
		opt.AccessKeyID,
		opt.SecretAccessKey,
		opt.BucketName,
		opt.BucketPath,
	)
	// end::mainfetch[]
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fetch Failure: %s\n", err)
		os.Exit(0)
	}

	// tag::mainprint[]
	fmt.Printf("Using Configuration: %s\n", globalConfig)
	// end::mainprint[]
}
