-- Restart-safe because MySQL DDL commits independently of migration bookkeeping.
SET @ddl = IF((SELECT COUNT(*) FROM information_schema.columns WHERE table_schema=DATABASE() AND table_name='exchange_posts' AND column_name='tags')=0, 'ALTER TABLE exchange_posts ADD COLUMN tags JSON NULL', 'SELECT 1');
PREPARE migration_stmt FROM @ddl;
EXECUTE migration_stmt;
DEALLOCATE PREPARE migration_stmt;
UPDATE exchange_posts SET tags=JSON_ARRAY() WHERE tags IS NULL;
ALTER TABLE exchange_posts MODIFY COLUMN tags JSON NOT NULL;
SET @ddl = IF((SELECT COUNT(*) FROM information_schema.columns WHERE table_schema=DATABASE() AND table_name='exchange_posts' AND column_name='kind')>0, 'ALTER TABLE exchange_posts DROP COLUMN kind', 'SELECT 1');
PREPARE migration_stmt FROM @ddl;
EXECUTE migration_stmt;
DEALLOCATE PREPARE migration_stmt;
