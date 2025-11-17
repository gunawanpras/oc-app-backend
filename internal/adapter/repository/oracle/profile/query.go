package oracle

var (
	queryGetProfile = `
		SELECT
			id,
			name,			
			TO_CLOB(UTL_RAW.CAST_TO_VARCHAR2(DBMS_LOB.SUBSTR(json_data, 4000, 1))) AS detail,
			created_at
		FROM myapp.profiles
	`

	queryCreateProfile = `
		INSERT INTO myapp.profiles (name, json_data)
		VALUES (:1, :2)
		RETURNING id INTO :3
	`

	queryUpdateProfile = `
		UPDATE myapp.profiles
		SET name = :1, 
			json_data = JSON_MERGEPATCH(json_data, :2)
		WHERE id = :3
		RETURNING id INTO :4
	`

	queryDeleteProfile = `
		DELETE FROM myapp.profiles WHERE id = :1
	`

	queryGetProfileViaProcedure = `BEGIN myapp.get_all_profiles(:1); END;`
)
