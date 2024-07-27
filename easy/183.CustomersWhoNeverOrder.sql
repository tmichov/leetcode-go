SELECT name from Customers
where id not in (select Orders.customerId from Orders);

SELECT name from Customers c
LEFT OUTER JOIN Orders o
ON c.id = o.customerId
WHERE o.customerId is NULL;
