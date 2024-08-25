USE TICKETMANAGER
GO

ALTER PROCEDURE CREATE_ACCOUNT (
	@name varchar(100),
	@email varchar(100),
	@password varchar(100),
	@departmentID int
) AS
BEGIN
BEGIN TRANSACTION
	BEGIN TRY 
		DECLARE @ID int
		SET @ID = (SELECT COUNT(A.ID_Account) FROM ACCOUNTS A) + 1

		DECLARE @IsPM bit
		SET @IsPM = 0
		IF @departmentID = 6 or @departmentID = 7
		BEGIN
			SET @IsPM = 1
		END

		INSERT INTO ACCOUNTS (ID_Account, Name, Email, Passwd, IsPM, Department_FK, IsStarted)
		VALUES (@ID, @name, @email, @password, @IsPM, @departmentID, 0)

		SELECT A.ID_Account
		FROM ACCOUNTS A
		WHERE A.ID_Account = @ID


		COMMIT TRANSACTION;
	END TRY
	BEGIN CATCH
		IF XACT_STATE() <> 0
		BEGIN
			ROLLBACK TRANSACTION;

		END
		DECLARE @ErrorMessage NVARCHAR(4000)
		DECLARE @ErrorSeverity INT
		DECLARE @ErrorState INT
		SELECT 
			@ErrorMessage = ERROR_MESSAGE(),
			@ErrorSeverity = ERROR_SEVERITY(),
			@ErrorState = ERROR_STATE();
		RAISERROR (@ErrorMessage, @ErrorSeverity, @ErrorState)
	END CATCH
		
END