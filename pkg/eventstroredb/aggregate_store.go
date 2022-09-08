package eventstroredb

import (
	"context"
	"fmt"
	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/ahmetb/go-linq/v3"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/core"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/core/domain"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/es"
	appendResult "github.com/mehdihadeli/store-golang-microservice-sample/pkg/es/append_result"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/es/contracts/store"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/es/stream_name"
	readPosition "github.com/mehdihadeli/store-golang-microservice-sample/pkg/es/stream_position/read_position"
	expectedStreamVersion "github.com/mehdihadeli/store-golang-microservice-sample/pkg/es/stream_version"
	esErrors "github.com/mehdihadeli/store-golang-microservice-sample/pkg/eventstroredb/errors"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/logger"
	typeMapper "github.com/mehdihadeli/store-golang-microservice-sample/pkg/reflection/type_mappper"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/serializer/jsonSerializer"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"reflect"
)

type esdbAggregateStore[T es.IHaveEventSourcedAggregate] struct {
	log        logger.Logger
	eventStore store.EventStore
	serializer *EsdbSerializer
}

func NewEventStoreAggregateStore[T es.IHaveEventSourcedAggregate](log logger.Logger, eventStore store.EventStore, serializer *EsdbSerializer) *esdbAggregateStore[T] {
	return &esdbAggregateStore[T]{log: log, eventStore: eventStore, serializer: serializer}
}

func (a *esdbAggregateStore[T]) StoreWithVersion(aggregate T, metadata *core.Metadata, expectedVersion expectedStreamVersion.ExpectedStreamVersion, ctx context.Context) (*appendResult.AppendEventsResult, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "esdbAggregateStore.StoreWithVersion")
	defer span.Finish()
	span.LogFields(log.Object("Aggregate", aggregate))
	span.LogFields(log.String("AggregateID", aggregate.Id().String()))

	if len(aggregate.UncommittedEvents()) == 0 {
		a.log.Infow(fmt.Sprintf("[esdbAggregateStore.StoreWithVersion] No events to store for aggregateId %s", aggregate.Id()), logger.Fields{"AggregateID": aggregate.Id()})
		return appendResult.NoOp, nil
	}

	streamId := streamName.For[T](aggregate)
	span.LogFields(log.String("StreamId", streamId.String()))

	var streamEvents []*es.StreamEvent

	linq.From(aggregate.UncommittedEvents()).SelectIndexedT(func(i int, domainEvent domain.IDomainEvent) *es.StreamEvent {
		var inInterface map[string]interface{}
		err := jsonSerializer.DecodeWithMapStructure(domainEvent, &inInterface)
		if err != nil {
			return nil
		}
		return a.serializer.DomainEventToStreamEvent(domainEvent, metadata, int64(i)+aggregate.OriginalVersion())
	}).ToSlice(&streamEvents)

	streamAppendResult, err := a.eventStore.AppendEvents(streamId, expectedVersion, streamEvents, ctx)
	if err != nil {
		return nil, tracing.TraceWithErr(span, errors.Wrapf(err, "[esdbAggregateStore_StoreWithVersion:AppendEvents] error in storing aggregate with id {%d}", aggregate.Id()))
	}

	aggregate.MarkUncommittedEventAsCommitted()

	span.LogFields(log.Object("StreamAppendResult", streamAppendResult))

	a.log.Infow(fmt.Sprintf("[esdbAggregateStore.StoreWithVersion] aggregate with id %d stored successfully", aggregate.Id()), logger.Fields{"AggregateID": aggregate.Id()})

	return streamAppendResult, nil
}

func (a *esdbAggregateStore[T]) Store(aggregate T, metadata *core.Metadata, ctx context.Context) (*appendResult.AppendEventsResult, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "esdbAggregateStore.Store")
	defer span.Finish()
	span.LogFields(log.Object("Aggregate", aggregate))
	span.LogFields(log.String("AggregateID", aggregate.Id().String()))

	expectedVersion := expectedStreamVersion.FromInt64(aggregate.OriginalVersion())

	streamAppendResult, err := a.StoreWithVersion(aggregate, metadata, expectedVersion, ctx)
	if err != nil {
		return nil, tracing.TraceWithErr(span, errors.Wrapf(err, "[esdbAggregateStore_Store:StoreWithVersion] failed to store aggregate with id{%v}", aggregate.Id()))
	}

	span.LogFields(log.Object("StreamAppendResult", streamAppendResult))

	a.log.Infow(fmt.Sprintf("[esdbAggregateStore.Store] aggregate with id %d stored successfully", aggregate.Id()), logger.Fields{"AggregateID": aggregate.Id()})

	return streamAppendResult, nil
}

func (a *esdbAggregateStore[T]) Load(ctx context.Context, aggregateId uuid.UUID) (T, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "esdbAggregateStore.Load")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", aggregateId.String()))

	position := readPosition.Start

	return a.LoadWithReadPosition(ctx, aggregateId, position)
}

func (a *esdbAggregateStore[T]) LoadWithReadPosition(ctx context.Context, aggregateId uuid.UUID, position readPosition.StreamReadPosition) (T, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "esdbAggregateStore.LoadWithReadPosition")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", aggregateId.String()))

	var typeNameType T
	aggregateInstance := typeMapper.InstancePointerByTypeName(typeMapper.GetTypeName(typeNameType))
	aggregate, ok := aggregateInstance.(T)
	if !ok {
		return *new(T), errors.New(fmt.Sprintf("[esdbAggregateStore_LoadWithReadPosition] aggregate is not a %s", typeMapper.GetTypeName(typeNameType)))
	}

	method := reflect.ValueOf(aggregate).MethodByName("NewEmptyAggregate")
	if !method.IsValid() {
		return *new(T), errors.New("[esdbAggregateStore_LoadWithReadPosition:MethodByName] aggregate does not have a `NewEmptyAggregate` method")
	}

	method.Call([]reflect.Value{})

	streamId := streamName.ForID[T](aggregateId)
	span.LogFields(log.String("StreamId", streamId.String()))

	streamEvents, err := a.getStreamEvents(streamId, position, ctx)
	if errors.Is(err, esdb.ErrStreamNotFound) || len(streamEvents) == 0 {
		return *new(T), tracing.TraceWithErr(span, errors.WithMessage(esErrors.NewAggregateNotFoundError(err, aggregateId), "[esdbAggregateStore.LoadWithReadPosition] error in loading aggregate"))
	}
	if err != nil {
		return *new(T), tracing.TraceWithErr(span, errors.Wrapf(err, "[esdbAggregateStore.LoadWithReadPosition:MethodByName] error in loading aggregate {%s}", aggregateId.String()))
	}

	var metadata *core.Metadata
	var domainEvents []domain.IDomainEvent

	linq.From(streamEvents).Distinct().SelectT(func(streamEvent *es.StreamEvent) domain.IDomainEvent {
		metadata = streamEvent.Metadata
		return streamEvent.Event
	}).ToSlice(&domainEvents)

	err = aggregate.LoadFromHistory(domainEvents, metadata)
	if err != nil {
		return *new(T), tracing.TraceWithErr(span, err)
	}

	a.log.Infow(fmt.Sprintf("Loaded aggregate with streamId {%s} and aggregateId {%s}",
		streamId.String(),
		aggregateId.String()),
		logger.Fields{"AggregateID": aggregateId.String(), "StreamId": streamId.String()})

	span.LogFields(log.Object("Aggregate", aggregate))

	return aggregate, nil
}

func (a *esdbAggregateStore[T]) Exists(ctx context.Context, aggregateId uuid.UUID) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "esdbAggregateStore.Exists")
	defer span.Finish()
	span.LogFields(log.String("AggregateId", aggregateId.String()))

	streamId := streamName.ForID[T](aggregateId)
	span.LogFields(log.String("StreamId", streamId.String()))

	return a.eventStore.StreamExists(streamId, ctx)
}

func (a *esdbAggregateStore[T]) getStreamEvents(streamId streamName.StreamName, position readPosition.StreamReadPosition, ctx context.Context) ([]*es.StreamEvent, error) {
	pageSize := 500
	var streamEvents []*es.StreamEvent

	for true {
		events, err := a.eventStore.ReadEvents(streamId, position, uint64(pageSize), ctx)
		if err != nil {
			return nil, errors.Wrapf(err, "[esdbAggregateStore_getStreamEvents:ReadEvents] failed to read events")
		}
		streamEvents = append(streamEvents, events...)
		if len(events) < pageSize {
			break
		}
		position = readPosition.FromInt64(int64(len(events)) + position.Value())
	}

	return streamEvents, nil
}
