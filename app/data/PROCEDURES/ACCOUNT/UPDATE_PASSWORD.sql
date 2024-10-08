ALTER PROCEDURE UPDATE_ACCOUNT_PASSWORD (
		@ID int,
		@newPassword varchar(100)
	) AS
	BEGIN 
	BEGIN TRANSACTION 
		BEGIN TRY
			UPDATE ACCOUNTS 
			SET ACCOUNTS.Passwd = @newPassword
			WHERE ACCOUNTS.ID_Account = @ID

			UPDATE ACCOUNTS 
			SET ACCOUNTS.IsStarted = 1
			WHERE ACCOUNTS.ID_Account = @ID

			SELECT A.ID_Account
			FROM ACCOUNTS A
			WHERE A.ID_Account = @ID

			COMMIT TRANSACTION
			RETURN
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
	