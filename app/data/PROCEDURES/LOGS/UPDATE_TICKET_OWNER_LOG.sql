CREATE PROCEDURE UPDATE_TICKET_OWNER_LOG (
	@ID_Ticket int,
	@OldOwner varchar(100),
	@NewOwner varchar(100),
	@date datetime
) AS 
BEGIN 
BEGIN TRANSACTION 
	BEGIN TRY
		DECLARE @ID_LOG int
		SET @ID_LOG = (SELECT COUNT(L.ID_Log) FROM LOGS L) + 1

		DECLARE @LogDesc varchar(500)
		SET @LogDesc = CONCAT(@date, '|| Ticket #', @ID_Ticket,' || Assigning a new owner || Old Owner: ',@OldOwner,' || New Owner: ', @NewOwner ,' ||')

		INSERT INTO LOGS
		VALUES (
			@ID_Ticket, @ID_LOG, @date, @LogDesc 
			)
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


