ALTER PROCEDURE UPDATE_TICKET_STATUS (
	@ID int,
	@ID_Status int,
	@date datetime
) AS
BEGIN
BEGIN TRANSACTION
	BEGIN TRY 
		DECLARE @OldStatusName varchar(50)
		DECLARE @NewStatusName varchar(50)

		SET @OldStatusName = (
			SELECT TS.Name
			FROM T_STATUS TS
			JOIN TICKETS T ON TS.ID_Status = T.Status_FK
			WHERE T.ID_Ticket = @ID
		)



		SET @NewStatusName = (
			SELECT TS.Name 
			FROM T_STATUS TS
			WHERE TS.ID_Status = @ID_Status
		)
		
		UPDATE TICKETS
		SET TICKETS.Status_FK = @ID_Status
		WHERE TICKETS.ID_Ticket = @ID

		EXEC UPDATE_TICKET_STATUS_LOG @ID, @OldStatusName, @NewStatusName, @date

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
