ALTER PROCEDURE CREATE_ACCOUNT (
	@name varchar(100),
	@email varchar(100),
	@departmentID int,
	@IsPM int 
) AS
BEGIN
BEGIN TRANSACTION
	BEGIN TRY 
		DECLARE @ID int
			SET @ID = (SELECT COUNT(A.ID_Account) FROM ACCOUNTS A) + 1
			INSERT INTO ACCOUNTS 
			VALUES (@ID, @name, @email, 'R3set3@able_P@ssw0rd_4312', @IsPM, @departmentID, 0)
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