CREATE DATABASE IF NOT EXISTS voicewriter CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE voicewriter;

-- 场景表
CREATE TABLE IF NOT EXISTS scenes (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL COMMENT '场景名称',
    description TEXT COMMENT '场景描述',
    icon VARCHAR(50) COMMENT '图标',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='场景表';

-- 句子表
CREATE TABLE IF NOT EXISTS sentences (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    scene_id INT UNSIGNED NOT NULL COMMENT '场景ID',
    content TEXT NOT NULL COMMENT '英文句子',
    translation TEXT COMMENT '中文翻译',
    audio_url VARCHAR(255) COMMENT '音频URL',
    difficulty VARCHAR(20) DEFAULT 'easy' COMMENT '难度: easy, medium, hard',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
    INDEX idx_scene_id (scene_id),
    INDEX idx_deleted_at (deleted_at),
    FOREIGN KEY (scene_id) REFERENCES scenes(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='句子表';

-- 用户进度表
CREATE TABLE IF NOT EXISTS user_progress (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id VARCHAR(100) NOT NULL COMMENT '用户ID',
    sentence_id INT UNSIGNED NOT NULL COMMENT '句子ID',
    completed BOOLEAN DEFAULT FALSE COMMENT '是否完成',
    attempts INT DEFAULT 0 COMMENT '尝试次数',
    last_attempt TIMESTAMP NULL DEFAULT NULL COMMENT '最后尝试时间',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间',
    INDEX idx_user_id (user_id),
    INDEX idx_sentence_id (sentence_id),
    INDEX idx_deleted_at (deleted_at),
    UNIQUE KEY uk_user_sentence (user_id, sentence_id, deleted_at),
    FOREIGN KEY (sentence_id) REFERENCES sentences(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户进度表';
