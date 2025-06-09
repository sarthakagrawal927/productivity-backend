package dao

var UpdateHabitFromLogs = `
UPDATE
	habits
SET
	existing_usage = COALESCE(hls.total_count, 0)
FROM (
	SELECT
		hl.habit_id,
		SUM(hl.count) AS total_count
	FROM
		habit_logs hl
		JOIN habits ON hl.habit_id = habits.id
	WHERE
		hl.habit_id = ? AND
		hl.logged_for_date > CURRENT_DATE - CASE WHEN habits.frequency_type = 1 THEN
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
	habits.id = hls.habit_id;`
