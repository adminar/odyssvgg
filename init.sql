-- 初始化数据库
CREATE DATABASE IF NOT EXISTS odyssey DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE odyssey;

-- 下面是表结构创建（无改动，只整理）

-- Table structure for table `build_records`
DROP TABLE IF EXISTS `build_records`;
CREATE TABLE `build_records` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `project_name` varchar(255) DEFAULT NULL,
  `env` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `build_user` varchar(255) DEFAULT NULL,
  `tag` varchar(255) DEFAULT NULL,
  `description` text,
  `status` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`,`project_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Table structure for table `build_service_records`
DROP TABLE IF EXISTS `build_service_records`;
CREATE TABLE `build_service_records` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `service_name` varchar(255) NOT NULL,
  `project_name` varchar(255) DEFAULT NULL,
  `env` varchar(255) DEFAULT NULL,
  `image` varchar(255) DEFAULT NULL,
  `build_record_name` varchar(255) DEFAULT NULL,
  `build_url` text,
  `status` varchar(255) DEFAULT 'pending',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `branch` varchar(255) DEFAULT NULL,
  `build_id` bigint DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `service_name` (`service_name`,`build_record_name`),
  KEY `build_record_name` (`build_record_name`),
  CONSTRAINT `build_service_records_ibfk_1` FOREIGN KEY (`build_record_name`) REFERENCES `build_records` (`name`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Table structure for table `clusters`
DROP TABLE IF EXISTS `clusters`;
CREATE TABLE `clusters` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(255) NOT NULL COMMENT '集群名称（唯一标识）',
  `api_server` varchar(255) NOT NULL COMMENT 'K8s API Server 地址',
  `config` text COMMENT 'Base64 或加密后的 kubeconfig 内容（可选）',
  `region` varchar(255) DEFAULT NULL COMMENT '集群所属区域（如 cn-hz）',
  `version` varchar(255) DEFAULT NULL COMMENT 'Kubernetes 版本号',
  `description` varchar(255) DEFAULT NULL COMMENT '备注说明',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Kubernetes 集群信息表';

-- Table structure for table `code_libraries`
DROP TABLE IF EXISTS `code_libraries`;
CREATE TABLE `code_libraries` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  `code_source_name` varchar(255) DEFAULT NULL,
  `project_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `codelibrary` (`name`),
  KEY `project` (`project_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Table structure for table `code_sources`
DROP TABLE IF EXISTS `code_sources`;
CREATE TABLE `code_sources` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  `private_token` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Table structure for table `deploy_records`
DROP TABLE IF EXISTS `deploy_records`;
CREATE TABLE `deploy_records` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `project_name` varchar(255) DEFAULT NULL,
  `build_record_name` varchar(255) DEFAULT NULL,
  `env` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `deploy_user` varchar(255) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `tag` varchar(255) DEFAULT NULL,
  `description` text,
  `cluster_names` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `build_record_name` (`build_record_name`),
  KEY `name` (`name`),
  CONSTRAINT `deploy_records_ibfk_1` FOREIGN KEY (`build_record_name`) REFERENCES `build_records` (`name`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Table structure for table `deploy_service_records`
DROP TABLE IF EXISTS `deploy_service_records`;
CREATE TABLE `deploy_service_records` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `service_name` varchar(255) NOT NULL,
  `project_name` varchar(255) DEFAULT NULL,
  `deploy_record_name` varchar(255) DEFAULT NULL,
  `env` varchar(255) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `image` text,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `cluster_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `deploy_record_name` (`deploy_record_name`),
  CONSTRAINT `deploy_service_records_ibfk_1` FOREIGN KEY (`deploy_record_name`) REFERENCES `deploy_records` (`name`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Table structure for table `envs`
DROP TABLE IF EXISTS `envs`;
CREATE TABLE `envs` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `name` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  `project_name` varchar(255) DEFAULT NULL,
  `namespace` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`,`project_name`),
  KEY `project_name` (`project_name`),
  CONSTRAINT `envs_ibfk_1` FOREIGN KEY (`project_name`) REFERENCES `projects` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Table structure for table `projects`
DROP TABLE IF EXISTS `projects`;
CREATE TABLE `projects` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `env` json DEFAULT NULL,
  `clusters` json DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Table structure for table `services`
DROP TABLE IF EXISTS `services`;
CREATE TABLE `services` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `name` varchar(255) DEFAULT NULL,
  `project_name` varchar(255) DEFAULT NULL,
  `env_name` varchar(255) DEFAULT NULL,
  `code_library_name` varchar(255) DEFAULT NULL,
  `deploy_map` json DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`,`project_name`,`env_name`),
  KEY `project_name` (`project_name`),
  CONSTRAINT `services_ibfk_1` FOREIGN KEY (`project_name`) REFERENCES `projects` (`name`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
