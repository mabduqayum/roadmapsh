-- Create custom types
CREATE TYPE article_status AS ENUM ('draft', 'published', 'archived');

-- Users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    is_admin BOOLEAN DEFAULT FALSE,
    last_login TIMESTAMP WITH TIME ZONE,
    CONSTRAINT valid_email CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$')
);

-- Articles table
CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    author_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    published_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    status article_status DEFAULT 'draft',
    view_count INTEGER DEFAULT 0,
    CONSTRAINT title_length CHECK (char_length(title) >= 3)
);

-- Create a function to automatically update timestamps
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger to fire before insert or update on users
CREATE TRIGGER set_timestamp_users
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION update_timestamp();

-- Create trigger to fire before insert or update on articles
CREATE TRIGGER set_timestamp_articles
    BEFORE UPDATE ON articles
    FOR EACH ROW
    EXECUTE FUNCTION update_timestamp();

-- Function to generate a unique slug from a title
CREATE OR REPLACE FUNCTION generate_unique_slug(title TEXT, table_name TEXT, slug_column TEXT)
RETURNS TEXT AS $$
DECLARE
    base_slug TEXT;
    new_slug TEXT;
    slug_exists BOOLEAN;
    max_base_length CONSTANT INTEGER := 64;
    counter INTEGER := 1;
BEGIN
    base_slug := LOWER(title);
    base_slug := REGEXP_REPLACE(base_slug, '[^a-z0-9\s-]', '', 'g');
    base_slug := REGEXP_REPLACE(base_slug, '\s+', '-', 'g');
    base_slug := TRIM(BOTH '-' FROM base_slug);
    base_slug := LEFT(base_slug, max_base_length);

    -- Initial slug attempt
    new_slug := base_slug;

    -- Check if slug exists and append counter if it does
    LOOP
    EXECUTE format('SELECT EXISTS(SELECT 1 FROM %I WHERE %I = $1)', table_name, slug_column)
    INTO slug_exists
    USING new_slug;

        EXIT WHEN NOT slug_exists;

        counter := counter + 1;
        new_slug := base_slug || '-' || counter::TEXT;
    END LOOP;

    RETURN new_slug;
END;
$$ LANGUAGE plpgsql;

-- Function to automatically generate slug for new articles
CREATE OR REPLACE FUNCTION set_article_slug()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.slug IS NULL OR NEW.slug = '' THEN
        NEW.slug := generate_unique_slug(NEW.title, 'articles', 'slug');
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to automatically generate slug for new articles
CREATE TRIGGER article_slug_trigger
    BEFORE INSERT ON articles
    FOR EACH ROW
    EXECUTE FUNCTION set_article_slug();
