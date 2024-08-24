CREATE PROCEDURE GET_TICKETS_BY_OWNER (
	@ID_Owner int
) AS
BEGIN 
	SELECT T.ID_Ticket, T.Title, T.T_Description, T.Creation_Date, A.ID_Account, A.Name, D.Name, S.Name
	FROM TICKETS T
	JOIN T_STATUS S ON T.Status_FK = S.ID_Status
	JOIN ACCOUNTS A ON T.Owner_FK = A.ID_Account
	JOIN DEPARTMENTS D ON T.Department_FK = D.ID_Department
	WHERE T.Owner_FK = @ID_Owner
END