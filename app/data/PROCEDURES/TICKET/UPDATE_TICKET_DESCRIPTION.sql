ALTER PROCEDURE UPDATE_TICKET_DESCRIPTION (
	@ID int,
	@newDesc text,
	@date datetime
) AS
BEGIN
BEGIN TRANSACTION
	BEGIN TRY 
		UPDATE TICKETS
		SET TICKETS.T_Description = @newDesc
		WHERE TICKETS.ID_Ticket = @ID

		DECLARE @OwnerName varchar(100)
		SET @OwnerName = (
			SELECT A.Name
			FROM ACCOUNTS A
			JOIN TICKETS T ON A.ID_Account = T.Owner_FK
			WHERE T.ID_Ticket = @ID
		)

		EXEC UPDATE_TICKET_DESCRIPTION_LOG @ID, @OwnerName, @newDesc, @date

		COMMIT TRANSACTION

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
