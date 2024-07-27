SELECT r.name as Employee FROM Employee as r 
INNER JOIN Employee as m
ON r.managerId = m.id and r.salary > m.salary;
