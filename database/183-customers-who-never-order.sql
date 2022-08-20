-- https://leetcode.com/problems/customers-who-never-order/

SELECT name as Customers
FROM Customers as c LEFT OUTER JOIN Orders as o ON c.id = o.customerId
WHERE o.id IS NULL;