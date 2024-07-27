SELECT Person.firstName, Person.lastName, Address.city, Address.state FROM Person
LEFT OUTER JOIN Address
ON Address.personId = Person.personId;
