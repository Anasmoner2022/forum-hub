package lab

import (
	"context"
	"fmt"
)

// RecordDomainEvent: Insert event inside the domain transaction with stable event ID.
// TODO: define these fields from the specification and use trusted context.
type RecordDomainEventInput struct{}
type RecordDomainEventOutput struct{}

func RecordDomainEvent(ctx context.Context, input RecordDomainEventInput) (RecordDomainEventOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return RecordDomainEventOutput{}, fmt.Errorf("TODO: RecordDomainEvent")
}

// DispatchEvents: Deliver a bounded batch idempotently.
// TODO: define these fields from the specification and use trusted context.
type DispatchEventsInput struct{}
type DispatchEventsOutput struct{}

func DispatchEvents(ctx context.Context, input DispatchEventsInput) (DispatchEventsOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return DispatchEventsOutput{}, fmt.Errorf("TODO: DispatchEvents")
}

// ListInbox: Return only authorized recipient notifications.
// TODO: define these fields from the specification and use trusted context.
type ListInboxInput struct{}
type ListInboxOutput struct{}

func ListInbox(ctx context.Context, input ListInboxInput) (ListInboxOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return ListInboxOutput{}, fmt.Errorf("TODO: ListInbox")
}

// MarkNotificationRead: Update only that recipient’s notification.
// TODO: define these fields from the specification and use trusted context.
type MarkNotificationReadInput struct{}
type MarkNotificationReadOutput struct{}

func MarkNotificationRead(ctx context.Context, input MarkNotificationReadInput) (MarkNotificationReadOutput, error) {
	// TODO: implement the assessed behavior; keep signatures typed.
	return MarkNotificationReadOutput{}, fmt.Errorf("TODO: MarkNotificationRead")
}
