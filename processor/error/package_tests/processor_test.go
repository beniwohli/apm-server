package package_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	er "github.com/elastic/apm-server/processor/error"
	"github.com/elastic/apm-server/tests"
)

// ensure all valid documents pass through the whole validation and transformation process
func TestProcessorOK(t *testing.T) {
	requestInfo := []tests.RequestInfo{
		{Name: "TestProcessErrorMinimalPayloadException", Path: "data/valid/error/minimal_payload_exception.json"},
		{Name: "TestProcessErrorMinimalPayloadLog", Path: "data/valid/error/minimal_payload_log.json"},
		{Name: "TestProcessErrorMinimalService", Path: "data/valid/error/minimal_service.json"},
		{Name: "TestProcessErrorMinimalProcess", Path: "data/valid/error/minimal_process.json"},
		{Name: "TestProcessErrorFull", Path: "data/valid/error/payload.json"},
		{Name: "TestProcessErrorNullValues", Path: "data/valid/error/null_values.json"},
	}
	tests.TestProcessRequests(t, er.NewProcessor(), requestInfo, map[string]string{})
}

// ensure invalid documents fail the json schema validation already
func TestProcessorFailedValidation(t *testing.T) {
	data, err := tests.LoadInvalidData("error")
	assert.Nil(t, err)
	err = er.NewProcessor().Validate(data)
	assert.NotNil(t, err)
}
