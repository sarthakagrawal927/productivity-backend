package dao

var GetFoodConsumptionLogs = `
	SELECT
		sum(fc.quantity * fi.fiber) fiber,
		sum(fc.quantity * fi.kcal) kcal,
		sum(fc.quantity * fi.protein) protein,
		fc. "date"
	FROM
		food_consumptions fc
		LEFT JOIN food_items fi ON fc.food_item_id = fi.id
	GROUP BY
		fc. "date" ORDER BY fc."date" DESC`

var GetNutrientsConsumedForDate = `
SELECT
	name,
	kcal * quantity kcal,
	protein * quantity protein,
	fiber * quantity fiber,
	quantity,
	created_at
FROM
	food_consumptions fc
	LEFT JOIN food_items fi ON fc.food_item_id = fi.id
WHERE
	fc. "date" = ?`
