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