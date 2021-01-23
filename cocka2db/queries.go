package cocka2db

const (
	noteTest = "select count(*) from note "

	insertUser = "INSERT into user(email, password, web_enabled) " +
		" values(?,?,?)"

	updateUser = " UPDATE user SET password = ?, web_enabled = ? " +
		" WHERE email = ?"

	getUser = "SELECT email, password, web_enabled " +
		" FROM user " +
		" WHERE email = ? "

	deleteUser = "DELETE FROM user WHERE email = ? "

	insertNote = "INSERT into note(title, type, owner_email, last_used) " +
		" values(?,?,?,?)"

	updateNote = " UPDATE note SET title = ?, type = ?, last_used = ? " +
		" WHERE id = ?"

	getNote = "SELECT id, title, type, owner_email, last_used " +
		" FROM note " +
		" WHERE id = ? "

	getNoteList = " SELECT n.id, n.title, n.type, n.owner_email, n.last_used " +
		" FROM note n " +
		" inner join note_users nu " +
		" on n.id = nu.note_id " +
		" inner join user u " +
		" on u.email = nu.user_email " +
		" where nu.user_email = ? " +
		" order by n.title "

	deleteNote = "DELETE FROM note WHERE id = ? "

	insertNoteUser = "INSERT into note_users(note_id, user_email) " +
		" values(?,?)"

	getNoteUserList = "SELECT user_email " +
		" FROM note_users " +
		" WHERE note_id = ? and user_email != ? "

	deleteNoteUser = "DELETE FROM note_users WHERE note_id = ? and user_email = ? "

	insertCheckboxNoteItem = "INSERT into checkbox_note_item(text, checked, note_id) " +
		" values(?,?,?)"

	updateCheckboxNoteItem = " UPDATE checkbox_note_item SET text = ?, checked = ? " +
		" WHERE id = ?"

	getCheckboxNoteItemList = "SELECT id, text, checked, note_id " +
		" FROM checkbox_note_item " +
		" WHERE note_id = ? " +
		" order by text "

	deleteCheckboxNoteItem = "DELETE FROM checkbox_note_item WHERE id = ? "

	insertNoteItem = "INSERT into note_item(text, note_id) " +
		" values(?,?)"

	updateNoteItem = " UPDATE note_item SET text = ? " +
		" WHERE id = ?"

	getNoteItemList = "SELECT id, text, note_id " +
		" FROM note_item " +
		" WHERE note_id = ? " +
		" order by text "

	deleteNoteItem = "DELETE FROM note_item WHERE id = ? "

	getMailServer = "SELECT id, mail_server, secure_connection, port, debug, " +
		" username, password, send_address " +
		" FROM mail_server "
)
