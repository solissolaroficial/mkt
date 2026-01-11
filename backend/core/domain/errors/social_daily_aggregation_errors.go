package errors

import "errors"

var (
	// ErrSocialDailyAggregationNotFound is returned when a daily aggregation is not found
	ErrSocialDailyAggregationNotFound = errors.New("social daily aggregation not found")

	// ErrSocialDailyAggregationIDRequired is returned when ID is required
	ErrSocialDailyAggregationIDRequired = errors.New("social daily aggregation ID is required")

	// ErrInvalidSocialDailyAggregationBrandName is returned when brand name is invalid
	ErrInvalidSocialDailyAggregationBrandName = errors.New("invalid brand name")

	// ErrInvalidSocialDailyAggregationPlatform is returned when platform is invalid
	ErrInvalidSocialDailyAggregationPlatform = errors.New("invalid platform")

	// ErrInvalidSocialDailyAggregationDate is returned when date is invalid
	ErrInvalidSocialDailyAggregationDate = errors.New("invalid date")
)
