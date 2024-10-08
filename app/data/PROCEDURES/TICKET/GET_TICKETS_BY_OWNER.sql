ALTER PROCEDURE GET_TICKETS_BY_OWNER (
	@ID_Owner int
) AS
BEGIN 
	SELECT T.ID_Ticket, T.Title, T.T_Description, A.ID_Account, A.Name AS 'Owner', D.Name AS 'Department', S.Name as 'Status', T.Priority
	FROM TICKETS T
	JOIN T_STATUS S ON T.Status_FK = S.ID_Status
	JOIN ACCOUNTS A ON T.Owner_FK = A.ID_Account
	JOIN DEPARTMENTS D ON T.Department_FK = D.ID_Department
	WHERE T.Owner_FK = @ID_Owner
	ORDER BY T.Status_FK
END