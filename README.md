# arqbeam-requestsio
An Apache Beam sink for arqbeam-app using requests http.

This implementation uses carlmjohnson/requests to send a POST message to http api.

TL;DR

```go
package main

import (
	"context"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/textio"
	requestsio "github.com/arquivei/arqbeam-requestsio"
)

var (
	pipeline *beam.Pipeline
)

func main() {
	pipeline = getPipeline(context.Background())
}

type JSONExample struct {
	Data string `json:"data"`
}

func getPipeline(_ context.Context) *beam.Pipeline {
	if pipeline != nil {
		return pipeline
	}

	pipeline := beam.NewPipeline()
	s := pipeline.Root()

	// Read some files with textio default from apache beam go sdk
	readRows := textio.Read(s, "gs://apache-beam-samples/shakespeare/kinglear.txt")

	jsonData := beam.ParDo(s, func(elm string) JSONExample { return JSONExample{Data: elm} }, readRows)

	// Send each line to endpoint with requestsio from arqbeam-requestsio
	requestsio.Post(s, "localhost:8080", jsonData)

	return pipeline
}

```

Comments, discussions, issues and pull-requests are welcomed.
