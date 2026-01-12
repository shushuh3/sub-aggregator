-- name: CreateSubscription :one
INSERT INTO subscriptions (
	service_name,
	price,
	user_id,
	start_date,
	end_date
)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetSubscription :one
SELECT *
FROM subscriptions
WHERE id = $1;

-- name: UpdateSubscription :one
UPDATE subscriptions
SET
	service_name = $1,
	price = $2,
	start_date = $3,
	end_date = $4,
	updated_at = now()
WHERE id = $5
RETURNING *;

-- name: DeleteSubscription :one
DELETE FROM subscriptions
WHERE id = $1
RETURNING *;

-- name: ListSubscriptions :many
SELECT *
FROM subscriptions
WHERE (sqlc.narg(user_id)::uuid IS NULL OR user_id = sqlc.narg(user_id)::uuid)
	AND (sqlc.narg(service_name)::text IS NULL OR service_name ILIKE '%' || sqlc.narg(service_name)::text || '%')
ORDER BY created_at DESC
LIMIT sqlc.arg(page_limit) OFFSET sqlc.arg(page_offset); 

-- name: SumSubscriptionsCost :one
WITH bounds AS (
	SELECT
		date_trunc('month', sqlc.arg(period_start)::date) AS period_start,
		date_trunc('month', sqlc.arg(period_end)::date) AS period_end
),
active AS (
	SELECT
		s.price,
		GREATEST(date_trunc('month', s.start_date), b.period_start) AS overlap_start,
		LEAST(date_trunc('month', COALESCE(s.end_date, b.period_end)), b.period_end) AS overlap_end
	FROM subscriptions s
	CROSS JOIN bounds b
	WHERE s.start_date <= b.period_end
		AND (s.end_date IS NULL OR s.end_date >= b.period_start)
		AND (sqlc.narg(user_id)::uuid IS NULL OR s.user_id = sqlc.narg(user_id)::uuid)
		AND (sqlc.narg(service_name)::text IS NULL OR s.service_name ILIKE '%' || sqlc.narg(service_name)::text || '%')
)
SELECT COALESCE(
	SUM(
		price * (
			DATE_PART('year', age(overlap_end, overlap_start)) * 12 +
			DATE_PART('month', age(overlap_end, overlap_start)) +
			1
		)
	),
	0
)::bigint AS total_cost
FROM active;
