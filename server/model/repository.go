package model

import (
	"context"

	"github.com/kubeshop/tracetest/server/pkg/id"
	"go.opentelemetry.io/otel/trace"
)

type List[T any] struct {
	Items      []T
	TotalCount int
}

type TestRepository interface {
	CreateTest(context.Context, Test) (Test, error)
	UpdateTest(context.Context, Test) (Test, error)
	DeleteTest(context.Context, Test) error
	TestIDExists(context.Context, id.ID) (bool, error)
	GetLatestTestVersion(context.Context, id.ID) (Test, error)
	GetTestVersion(_ context.Context, _ id.ID, version int) (Test, error)
	GetTests(_ context.Context, take, skip int32, query, sortBy, sortDirection string) (List[Test], error)
}

type RunRepository interface {
	CreateRun(context.Context, Test, Run) (Run, error)
	UpdateRun(context.Context, Run) error
	DeleteRun(context.Context, Run) error
	GetRun(_ context.Context, testID id.ID, runID int) (Run, error)
	GetTestRuns(_ context.Context, _ Test, take, skip int32) (List[Run], error)
	GetRunByTraceID(context.Context, trace.TraceID) (Run, error)
	GetLatestRunByTestVersion(context.Context, id.ID, int) (Run, error)
}

type TestRunEventRepository interface {
	CreateTestRunEvent(context.Context, TestRunEvent) error
	GetTestRunEvents(context.Context, id.ID, int) ([]TestRunEvent, error)
}

type Repository interface {
	TestRepository
	RunRepository

	TestRunEventRepository

	ServerID() (id string, isNew bool, _ error)
	Close() error
	Drop() error
}
