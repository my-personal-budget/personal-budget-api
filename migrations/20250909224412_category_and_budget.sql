-- Write your migration here

CREATE TABLE IF NOT EXISTS Category (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Budget table with DATE for month
CREATE TABLE IF NOT EXISTS Budget (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    month DATE NOT NULL,             -- first day of month, e.g., '2025-01-01'
    category_id INTEGER NOT NULL,  -- FK to categories
    budget REAL NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories(id),
    UNIQUE(month, category_id)      -- prevent duplicates for same month/category
);
