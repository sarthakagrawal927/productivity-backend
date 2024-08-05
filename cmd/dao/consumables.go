package dao

var GetFoodConsumptionLogs = `
	WITH grouped_data AS (
		SELECT
			sum(fc.quantity * fi.fiber) AS fiber,
			sum(fc.quantity * fi.kcal) AS total_kcal,
			sum(fc.quantity * fi.protein) AS protein,
			sum(fc.quantity * fi.fat) AS fat,
			sum(fc.quantity * fi.carbs) AS carbs,
			%s AS date_group,
			COUNT(DISTINCT fc."date") AS num_days
		FROM
			food_consumptions fc
			LEFT JOIN food_items fi ON fc.food_item_id = fi.id
		WHERE fc.user_id = ?
		GROUP BY
			date_group
	)
	SELECT
		fiber / num_days AS fiber,
		total_kcal / num_days AS kcal,
		protein / num_days AS protein,
		fat / num_days AS fat,
		carbs / num_days AS carbs,
		date_group as date
	FROM
		grouped_data
	ORDER BY
		date DESC`

var GetNutrientsConsumedForDate = `
SELECT
	name,
	kcal * quantity kcal,
	protein * quantity protein,
	fiber * quantity fiber,
	fat * quantity fat,
	carbs * quantity carbs,
	quantity,
	created_at
FROM
	food_consumptions fc
	LEFT JOIN food_items fi ON fc.food_item_id = fi.id
WHERE
	fc."date" = ? AND fc.user_id = ?`
