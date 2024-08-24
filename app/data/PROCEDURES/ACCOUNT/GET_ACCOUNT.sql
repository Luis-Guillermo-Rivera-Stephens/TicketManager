CREATE PROCEDURE GET_ACCOUNT(
	@ID int
) AS
BEGIN 
	SELECT A.ID_Account, A.Name, A.Email, A.Passwd, A.IsStarted, A.Department_FK, A.IsPM
	FROM ACCOUNTS A
	WHERE A.ID_Account = @ID
END
