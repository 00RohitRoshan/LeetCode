/* Write your PL/SQL query statement below */
Select firstName,lastName,city,state From Person Left Outer Join Address On (Person.personId=Address.personId)