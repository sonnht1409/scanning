CREATE TABLE `scannings` (
  `id` bigint(20) unsigned PRIMARY KEY AUTO_INCREMENT, 
  `repo_name` varchar(255) NOT NULL, 
  `repo_url` varchar(255) NOT NULL, 
  `finding` JSON, 
  `scan_unique_id` varchar(255) NOT NULL, 
  `status` int NOT NULL DEFAULT '0', 
  `queued_at` timestamp, 
  `scanning_at` timestamp, 
  `finished_at` timestamp, 
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, 
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, 
  UNIQUE KEY `unique_index_scan_unique_id` (`scan_unique_id`), 
  KEY `index_updated_at` (`updated_at`), 
  KEY `index_repo_name` (`repo_name`)
) ENGINE = InnoDB  DEFAULT CHARSET = utf8;