package db

var InitStatements = []string{
	"CREATE SCHEMA IF NOT EXISTS \"contactsApi\"",
	`CREATE TABLE IF NOT EXISTS "contactsApi".users(
		id SERIAL,
		"firstName" varchar(90) DEFAULT NULL,
		"lastName" varchar(90) DEFAULT NULL,
		updated_at timestamp DEFAULT NOW(),
		created_at timestamp DEFAULT NOW(),
		CONSTRAINT pk_users_id PRIMARY KEY (id) 
	);`,
	` CREATE UNIQUE INDEX IF NOT EXISTS pk_users_index ON "contactsApi".users
	USING btree
	(
	  id ASC NULLS LAST
	);`,
	`CREATE INDEX IF NOT EXISTS pk_users_created_at ON "contactsApi".users
	USING btree
	(
	  created_at ASC NULLS LAST
	);`,
	`CREATE INDEX IF NOT EXISTS pk_users_updated_at ON "contactsApi".users
	USING btree
	(
	  updated_at ASC NULLS LAST
	);`,
	`CREATE TABLE IF NOT EXISTS "contactsApi".contacts(
		id SERIAL,
		user_id bigint NOT NULL,
		"firstName" varchar(90) DEFAULT NULL,
		"lastName" varchar(90) DEFAULT NULL,
		"email" varchar(90) DEFAULT NULL,
		"phone" varchar(90) DEFAULT NULL,
		updated_at timestamp DEFAULT NOW(),
		created_at timestamp DEFAULT NOW(),
		CONSTRAINT pk_contacts_id PRIMARY KEY (id) 
	);`,
	` CREATE UNIQUE INDEX IF NOT EXISTS pk_contacts_index ON "contactsApi".contacts
	USING btree
	(
	  id ASC NULLS LAST
	);`,
	`CREATE INDEX IF NOT EXISTS pk_contacts_created_at ON "contactsApi".contacts
	USING btree
	(
	  created_at ASC NULLS LAST
	);`,
	`CREATE INDEX IF NOT EXISTS pk_contacts_created_at ON "contactsApi".contacts
	USING btree
	(
	  created_at ASC NULLS LAST
	);`,
	`CREATE INDEX IF NOT EXISTS pk_contacts_updated_at ON "contactsApi".contacts
	USING btree
	(
	  updated_at ASC NULLS LAST
	);`,
	`create or replace function create_constraint_if_not_exists (
    s_name text, t_name text, c_name text, constraint_sql text
) 
returns void AS
$$
begin
    -- Look for our constraint
    if not exists (select constraint_name, constraint_schema
                   from information_schema.constraint_column_usage 
                   where constraint_name = c_name and constraint_schema = s_name) then
        execute constraint_sql;
    end if;
end;
$$ language 'plpgsql'`,
	`SELECT create_constraint_if_not_exists(
        'contactsApi',
		'',
        'fk_users_user_id',
        'ALTER TABLE "contactsApi".contacts ADD CONSTRAINT fk_users_user_id FOREIGN KEY ("user_id")
REFERENCES "contactsApi".users (id) MATCH FULL
ON DELETE CASCADE ON UPDATE NO ACTION;')`,
	`CREATE OR REPLACE FUNCTION "contactsApi".set_timestamp()
RETURNS TRIGGER LANGUAGE 'plpgsql' AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$;`,
	`DROP TRIGGER IF EXISTS set_contacts_timestamp ON "contactsApi".contacts CASCADE`,
	`DROP TRIGGER IF EXISTS set_users_timestamp ON "contactsApi".users CASCADE`,
	`CREATE TRIGGER set_contacts_timestamp
	BEFORE UPDATE
	ON "contactsApi".contacts
	FOR EACH ROW
	EXECUTE PROCEDURE "contactsApi".set_timestamp();`,
	`CREATE TRIGGER set_users_timestamp
	BEFORE UPDATE
	ON "contactsApi".users
	FOR EACH ROW
	EXECUTE PROCEDURE "contactsApi".set_timestamp();`,
}
