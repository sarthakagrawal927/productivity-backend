# Backend

[Endpoints](https://galactic-escape-804413.postman.co/workspace/Stumble~8b18c535-4c33-4445-9f41-3fd645691d7d/collection/20124508-30a6e9dc-c460-473e-a0c3-e2fc2c5e066d?action=share&creator=20124508&active-environment=26183107-3fa6a988-04b0-403c-8575-7ec046deff9c)

## Notes

- Gorm adds deletedAt column in each table, and it always soft deletes. That's why there are no delete status in the models. Though it is to promote soft delete I feel having extra column can be avoided if I already have a status column in most tables which can be purposed into delete. Will do that after the project is in reasonable shape.

- Decided to maintain habit streak status in habit itself. Because if not done that, it results in complex query like:
```sql
SELECT
	*
FROM
	habits h
	LEFT JOIN (
		SELECT
			habit_id,
			sum(result_count)
		FROM
			habit_logs
		WHERE
			result_date >= CURRENT_DATE - CASE WHEN h.frequency_type = 1 THEN
				CURRENT_DATE - INTERVAL '1 days'
			WHEN h.frequency_type = 2 THEN
				CURRENT_DATE - INTERVAL '7 days'
			WHEN h.frequency_type = 3 THEN
				CURRENT_DATE - INTERVAL '1 months'
			ELSE
				CURRENT_DATE - INTERVAL '7 days'
			END
		GROUP BY
			habit_id) hls ON h.id = hls.habit_id
WHERE
	h.anti = FALSE;
```
(Query is not working, just to show the complexity)

- Also need to consider if in future we want to show performance of last 30 days whether to fetch via logs or maintain in habit itself.

To update the habit with logs:
(Can setup a cron or to fix old data, streak will be gone yes)

```sql
UPDATE
	habits
SET
	existing_usage = COALESCE(hls.total_result_count, 0),
	current_streak = COALESCE(current_streak, 0) + (
		CASE WHEN anti THEN
			CASE WHEN COALESCE(hls.total_result_count, 0) <= target THEN
				1
			ELSE
				0
			END
		ELSE
			CASE WHEN COALESCE(hls.total_result_count, 0) >= target THEN
				1
			ELSE
				0
			END
		END)
FROM (
	SELECT
		hl.habit_id,
		SUM(hl.result_count) AS total_result_count
	FROM
		habit_logs hl
		JOIN habits ON hl.habit_id = habits.id
	WHERE
		hl.result_date > CURRENT_DATE - CASE WHEN habits.frequency_type = 1 THEN
			INTERVAL '1 days'
		WHEN habits.frequency_type = 2 THEN
			INTERVAL '7 days'
		WHEN habits.frequency_type = 3 THEN
			INTERVAL '1 months'
		ELSE
			INTERVAL '7 days'
		END
	GROUP BY
		hl.habit_id) hls
WHERE
	habits.id = hls.habit_id;
```

Because of complexity, currently considering to stop maintaining maxStreak. Will consider some more novel solution for it later on.