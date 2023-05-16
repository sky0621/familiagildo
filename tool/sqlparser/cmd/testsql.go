package main

const selectSQL01 = `
SELECT p.p_name, SUM(o.quantity), SUM(p.price * o.quantity)
FROM order_desc AS o RIGHT JOIN product AS p
ON p.p_id = o.p_id
GROUP BY p.p_id, p.p_name ORDER BY SUM(p.price * o.quantity) DESC;
`

const selectSQL02 = `
SELECT COUNT(bp.product_id) AS how_many_products,
COUNT(dev.account_id) AS how_many_developers,
COUNT(b.bug_id)/COUNT(dev.account_id) AS avg_bugs_per_developer,
COUNT(cust.account_id) AS how_many_customers
FROM Bugs b JOIN BugsProducts bp ON (b.bug_id = bp.bug_id)
JOIN Accounts dev ON (b.assigned_to = dev.account_id)
JOIN Accounts cust ON (b.reported_by = cust.account_id)
WHERE cust.email NOT LIKE '%@example.com'
GROUP BY bp.product_id;
`

const selectSQL03 = `
SELECT T.*
FROM (
 SELECT E.ID, E.NAME, B.NAME AS BRANCH_NAME, D.NAME AS DEPARTMENT_NAME
 FROM EMPLOYEES E
 JOIN BRANCH B ON E.BRANCH_ID = B.ID
 JOIN DEPARTMENT D ON E.DEPARTMENT_ID = D.ID
) T
WHERE T.ID = 1;
`

const insertSQL01 = `
INSERT INTO guild (name, status) VALUES ($1, 1)
RETURNING *;
`

const updateSQL01 = `
UPDATE guild SET status = 2 WHERE id = $1
RETURNING *;
`

const deleteSQL01 = `
DELETE FROM guest_token WHERE id = $1;
`

var sqls = [][]string{
	{"selectSQL01", selectSQL01},
	{"selectSQL02", selectSQL02},
	{"selectSQL03", selectSQL03},
	{"insertSQL01", insertSQL01},
	{"updateSQL01", updateSQL01},
	{"deleteSQL01", deleteSQL01},
}
