ALTER PROCEDURE UPDATE_TICKET_OWNER (
	@ID int,
	@ID_Owner INT,
	@date datetime
) AS
BEGIN
BEGIN TRANSACTION
	BEGIN TRY 

		DECLARE @OldOwnerName varchar(100)
		DECLARE @newOwnerName varchar(100)
		DECLARE @currentStatus int

		SET @OldOwnerName = (
			SELECT A.Name
			FROM ACCOUNTS A
			JOIN TICKETS T ON A.ID_Account = T.Owner_FK
			WHERE T.ID_Ticket = @ID
		)

		IF @OldOwnerName = NULL
		BEGIN
			SET @OldOwnerName = 'Unassigned'

		END

		SET @newOwnerName = (
			SELECT A.Name 
			FROM ACCOUNTS A
			WHERE A.ID_Account = @ID_Owner
		)

		SET @currentStatus = (
			SELECT  T.Status_FK
			FROM TICKETS T
			WHERE T.ID_Ticket = @ID
		)

		UPDATE TICKETS
		SET TICKETS.Owner_FK = @ID_Owner
		WHERE TICKETS.ID_Ticket = @ID

		EXEC UPDATE_TICKET_OWNER_LOG @ID, @OldOwnerName, @newOwnerName, @date

		IF @currentStatus = 1
		BEGIN
			EXEC UPDATE_TICKET_STATUS @ID, 2, @date
		END



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

