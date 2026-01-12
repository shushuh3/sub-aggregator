package domain

import (
	"context"

	"github.com/google/uuid"
)

// SubscriptionRepository defines storage operations for subscriptions.
type SubscriptionRepository interface {
	Create(ctx context.Context, params CreateSubscriptionParams) (Subscription, error)
	Get(ctx context.Context, id uuid.UUID) (Subscription, error)
	Update(ctx context.Context, params UpdateSubscriptionParams) (Subscription, error)
	Delete(ctx context.Context, id uuid.UUID) (Subscription, error)
	List(ctx context.Context, params ListSubscriptionsParams) ([]Subscription, error)
	SumCost(ctx context.Context, params SumCostParams) (int64, error)
}

type CreateSubscriptionParams struct {
	ServiceName string
	Price       uint32
	UserID      uuid.UUID
	StartDate   MonthYear
	EndDate     *MonthYear
}

type UpdateSubscriptionParams struct {
	ID          uuid.UUID
	ServiceName string
	Price       uint32
	StartDate   MonthYear
	EndDate     *MonthYear
}

type ListSubscriptionsParams struct {
	UserID      *uuid.UUID
	ServiceName *string
	Page        int32
	PageSize    int32
}

type SumCostParams struct {
	PeriodStart MonthYear
	PeriodEnd   MonthYear
	UserID      *uuid.UUID
	ServiceName *string
}
