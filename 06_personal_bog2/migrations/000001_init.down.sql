-- Drop triggers
DROP TRIGGER IF EXISTS article_slug_trigger ON articles;
DROP TRIGGER IF EXISTS set_timestamp_articles ON articles;
DROP TRIGGER IF EXISTS set_timestamp_users ON users;

-- Drop functions
DROP FUNCTION IF EXISTS set_article_slug();
DROP FUNCTION IF EXISTS generate_unique_slug(TEXT, TEXT, TEXT);
DROP FUNCTION IF EXISTS update_timestamp();

-- Drop tables
DROP TABLE IF EXISTS articles;
DROP TABLE IF EXISTS users;

-- Drop custom types
DROP TYPE IF EXISTS article_status;
