CREATE PROCEDURE GET_UNASSIGNED_TICKETS_BY_DEPARTMENT (
	@ID_Department int
) AS
BEGIN 
	SELECT T.ID_Ticket, T.Title, T.T_Description, T.Creation_Date, A.ID_Account, A.Name AS 'Owner', D.Name AS 'Department', S.Name as 'Status'
	FROM TICKETS T
	JOIN T_STATUS S ON T.Status_FK = S.ID_Status
	LEFT JOIN ACCOUNTS A ON T.Owner_FK = A.ID_Account
	JOIN DEPARTMENTS D ON T.Department_FK = D.ID_Department
	WHERE T.Department_FK = @ID_Department and T.Status_FK = 1
END