ALTER PROCEDURE GET_OPEN_TICKETS_BY_DEPARTMENT (
	@ID_Department int
) AS
BEGIN 
	SELECT T.ID_Ticket, T.Title, T.T_Description, A.ID_Account, A.Name AS 'Owner', D.Name AS 'Department', S.Name as 'Status', T.Priority
	FROM TICKETS T
	JOIN T_STATUS S ON T.Status_FK = S.ID_Status
	LEFT JOIN ACCOUNTS A ON T.Owner_FK = A.ID_Account
	JOIN DEPARTMENTS D ON T.Department_FK = D.ID_Department
	WHERE T.Department_FK = @ID_Department and T.Status_FK <> 6 and T.Status_FK <> 7
	ORDER BY T.Status_FK
END