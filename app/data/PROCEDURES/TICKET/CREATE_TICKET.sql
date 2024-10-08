ALTER PROCEDURE CREATE_TICKET (
	@title varchar(100),
	@description text,
	@Department_FK int,
	@date datetime,
	@priority varchar(50)
)	AS
BEGIN 
BEGIN TRANSACTION
	BEGIN TRY 
		DECLARE @ID int
			SET @ID = (SELECT COUNT(T.ID_Ticket) FROM TICKETS T) + 1
			INSERT INTO TICKETS 
			VALUES (@ID, @title, @description, 1, NULL, @Department_FK, @date, @priority)

			EXEC CREATE_TICKET_LOG @ID, @title, @date

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