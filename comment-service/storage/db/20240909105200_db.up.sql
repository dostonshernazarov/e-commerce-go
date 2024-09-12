CREATE TABLE comments (
    comment_id SERIAL PRIMARY KEY,           
    post_id INT NOT NULL,                  
    user_id INT NOT NULL,                   
    message TEXT NOT NULL,                   
    comment_like INT DEFAULT 0,              
    created_at TIMESTAMPTZ DEFAULT NOW(),    
    updated_at TIMESTAMPTZ DEFAULT NOW()     
);
