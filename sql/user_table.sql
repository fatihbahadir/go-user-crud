CREATE TABLE users (
    id TEXT PRIMARY KEY,          
    name TEXT NOT NULL,           
    surname TEXT NOT NULL,        
    email TEXT NOT NULL UNIQUE,   
    phone_number TEXT UNIQUE,    
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);