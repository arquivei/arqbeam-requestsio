package requestsio

import (
	"bytes"
	"context"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"github.com/carlmjohnson/requests"
)

func init() {
	register.DoFn2x1[context.Context, any, []byte](&RequestsIO{})
}

type PubsubResultError struct {
	MessageID    string
	ErrorMessage string
}

func Post(s beam.Scope, url string, targetData beam.PCollection) beam.PCollection {
	s = s.Scope("arqbeam.requestsIO.Post")

	return beam.ParDo(s, &RequestsIO{
		URL: url,
	}, targetData)
}

type RequestsIO struct {
	URL string
}

func (fn *RequestsIO) ProcessElement(ctx context.Context, value any) []byte {
	var stationResp bytes.Buffer
	err := requests.
		URL(fn.URL).
		BodyJSON(value).
		ToBytesBuffer(&stationResp).
		Fetch(ctx)

	if err != nil {
		log.Errorf(ctx, err.Error())
	}

	return stationResp.Bytes()
}
