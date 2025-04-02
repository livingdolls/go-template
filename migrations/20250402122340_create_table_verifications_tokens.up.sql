CREATE TABLE verification_tokens (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    token VARCHAR(512) NOT NULL,
    user_id CHAR(36) NOT NULL,
    token_type ENUM('email_verification', 'password_reset', 'api_access') NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    used_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    INDEX idx_verification_tokens_token (token),
    INDEX idx_verification_tokens_user_id (user_id),
    INDEX idx_verification_tokens_expires_at (expires_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;