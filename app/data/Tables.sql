/*
Tables without any relation
*/

CREATE TABLE ACCOUNTS (
	ID_Account int primary key Not Null,
	Name varchar(100) Not Null,
	Email varchar(100) Not Null,
	Passwd varchar(150) Not Null,
	IsPM bit Not Null,
	Department_FK int Not Null,
	IsStarted bit Not Null
)

CREATE TABLE DEPARTMENTS (
	ID_Department int primary key Not Null,
	Name varchar(100) Not Null,
	D_Description text Not Null
)


CREATE TABLE TICKETS (
	ID_Ticket int primary key Not Null,
	Title varchar(50) Not Null,
	T_Description text Not Null,
	Status_FK int Not Null,
	Owner_FK int Null,
	Department_FK int Not Null
)

CREATE TABLE T_STATUS (
	ID_Status int primary key Not Null,
	Name varchar(100) Not Null
)

CREATE TABLE LOGS (
	ID_Ticket int primary key Not Null,
	ID_Log int Not Null,
	Log_Date date not null,
	Log_Description text not null
)