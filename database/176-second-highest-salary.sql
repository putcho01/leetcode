-- https://leetcode.com/problems/second-highest-salary/

SELECT IFNULL(
  (SELECT salary
   FROM Employee 
   where salary < (select MAX(salary) FROM Employee) 
   order by salary desc limit 1), 
  null
) as SecondHighestSalary;