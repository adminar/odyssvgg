-- MySQL dump 10.13  Distrib 8.0.41, for macos15.2 (arm64)
--
-- Host: 172.25.0.94    Database: odyssey
-- ------------------------------------------------------
-- Server version	8.3.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE DATABASE IF NOT EXISTS odyssey DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE odyssey;

--
-- Table structure for table `build_records`
--

DROP TABLE IF EXISTS `build_records`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `build_records` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '记录唯一标识，自增',
  `name` varchar(255) NOT NULL COMMENT '构建记录名称',
  `project_name` varchar(255) DEFAULT NULL COMMENT '构建所属项目',
  `env` varchar(255) DEFAULT NULL COMMENT '构建的环境',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `build_user` varchar(255) DEFAULT NULL COMMENT '谁构建的',
  `tag` varchar(255) DEFAULT NULL COMMENT '构建的标签(测试, 预发, 上线)',
  `description` text COMMENT '构建记录的描述',
  `status` varchar(255) DEFAULT NULL COMMENT '构建状态(Success, Pending, Failed)',
  PRIMARY KEY (`id`) COMMENT '主键：记录唯一标识',
  UNIQUE KEY `name` (`name`,`project_name`) COMMENT '唯一索引：构建记录名称和项目名称组合唯一'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='构建记录表，存储每次构建的相关信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `build_records`
--

/*!40000 ALTER TABLE `build_records` DISABLE KEYS */;
/*!40000 ALTER TABLE `build_records` ENABLE KEYS */;

--
-- Table structure for table `build_service_records`
--

DROP TABLE IF EXISTS `build_service_records`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `build_service_records` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '记录唯一标识，自增',
  `service_name` varchar(255) NOT NULL COMMENT '服务名称',
  `project_name` varchar(255) DEFAULT NULL COMMENT '服务所属项目',
  `env` varchar(255) DEFAULT NULL COMMENT '服务构建环境',
  `image` varchar(255) DEFAULT NULL COMMENT '构建出来的镜像',
  `build_record_name` varchar(255) DEFAULT NULL COMMENT '服务所属的构建记录名称',
  `build_url` text COMMENT '构建的 Jenkins URL',
  `status` varchar(255) DEFAULT 'pending' COMMENT '服务当前构建状态(Success, Pending, Failed, Abort)',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `branch` varchar(255) DEFAULT NULL COMMENT '构建分支',
  `build_id` bigint DEFAULT '0' COMMENT 'Jenkins 构建 ID',
  PRIMARY KEY (`id`) COMMENT '主键：记录唯一标识',
  UNIQUE KEY `service_name` (`service_name`,`build_record_name`) COMMENT '唯一索引：服务名称和构建记录名称组合唯一',
  KEY `build_record_name` (`build_record_name`) COMMENT '索引：构建记录名称',
  CONSTRAINT `build_service_records_ibfk_1` FOREIGN KEY (`build_record_name`) REFERENCES `build_records` (`name`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='构建服务记录表，存储每次服务构建的相关信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `build_service_records`
--

/*!40000 ALTER TABLE `build_service_records` DISABLE KEYS */;
/*!40000 ALTER TABLE `build_service_records` ENABLE KEYS */;

--
-- Table structure for table `clusters`
--

DROP TABLE IF EXISTS `clusters`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
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
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `clusters`
--

/*!40000 ALTER TABLE `clusters` DISABLE KEYS */;
/*!40000 ALTER TABLE `clusters` ENABLE KEYS */;

--
-- Table structure for table `code_libraries`
--

DROP TABLE IF EXISTS `code_libraries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `code_libraries` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '记录唯一标识，自增',
  `name` varchar(255) DEFAULT NULL COMMENT '代码库的名称',
  `url` varchar(255) DEFAULT NULL COMMENT '代码库地址',
  `code_source_name` varchar(255) DEFAULT NULL COMMENT '代码源名称',
  `project_name` varchar(255) DEFAULT NULL COMMENT '代码库所属项目',
  PRIMARY KEY (`id`),
  KEY `codelibrary` (`name`),
  KEY `project` (`project_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='代码库表，存储代码库的相关信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `code_libraries`
--

/*!40000 ALTER TABLE `code_libraries` DISABLE KEYS */;
/*!40000 ALTER TABLE `code_libraries` ENABLE KEYS */;

--
-- Table structure for table `code_sources`
--

DROP TABLE IF EXISTS `code_sources`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `code_sources` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL COMMENT '代码源名称',
  `type` varchar(255) DEFAULT NULL COMMENT '代码源类型(gitlab,github)',
  `url` varchar(255) DEFAULT NULL COMMENT '代码源地址',
  `private_token` varchar(255) DEFAULT NULL COMMENT '代码源私钥',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `code_sources`
--

/*!40000 ALTER TABLE `code_sources` DISABLE KEYS */;
/*!40000 ALTER TABLE `code_sources` ENABLE KEYS */;

--
-- Table structure for table `deploy_records`
--

DROP TABLE IF EXISTS `deploy_records`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `deploy_records` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '记录唯一标识，自增',
  `name` varchar(255) NOT NULL COMMENT '发布记录名称',
  `project_name` varchar(255) DEFAULT NULL COMMENT '发布记录所属项目',
  `build_record_name` varchar(255) DEFAULT NULL COMMENT '发布记录所属构建记录',
  `env` varchar(255) DEFAULT NULL COMMENT '发布记录所属环境',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `deploy_user` varchar(255) DEFAULT NULL COMMENT '谁发布的',
  `status` varchar(255) DEFAULT NULL COMMENT '发布状态',
  `tag` varchar(255) DEFAULT NULL COMMENT '发布标签(测试, 预发, 上线)',
  `description` text COMMENT '发布记录的描述',
  `cluster_names` varchar(255) DEFAULT NULL COMMENT '发布到哪些集群',
  PRIMARY KEY (`id`) COMMENT '主键：记录唯一标识',
  KEY `build_record_name` (`build_record_name`) COMMENT '索引：构建记录名称',
  KEY `name` (`name`) COMMENT '索引：发布记录名称',
  CONSTRAINT `deploy_records_ibfk_1` FOREIGN KEY (`build_record_name`) REFERENCES `build_records` (`name`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='发布记录表，存储每次发布的相关信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `deploy_records`
--

/*!40000 ALTER TABLE `deploy_records` DISABLE KEYS */;
/*!40000 ALTER TABLE `deploy_records` ENABLE KEYS */;

--
-- Table structure for table `deploy_service_records`
--

DROP TABLE IF EXISTS `deploy_service_records`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `deploy_service_records` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '记录唯一标识，自增',
  `service_name` varchar(255) NOT NULL COMMENT '服务名称',
  `project_name` varchar(255) DEFAULT NULL COMMENT '项目名称',
  `deploy_record_name` varchar(255) DEFAULT NULL COMMENT '部署记录名称',
  `env` varchar(255) DEFAULT NULL COMMENT '部署环境',
  `status` varchar(255) DEFAULT NULL COMMENT '部署状态',
  `image` text COMMENT '服务镜像信息',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  `cluster_name` varchar(255) DEFAULT NULL COMMENT '集群名称',
  PRIMARY KEY (`id`) COMMENT '主键：记录唯一标识',
  KEY `deploy_record_name` (`deploy_record_name`) COMMENT '索引：部署记录名称',
  CONSTRAINT `deploy_service_records_ibfk_1` FOREIGN KEY (`deploy_record_name`) REFERENCES `deploy_records` (`name`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='部署服务记录表，存储每次部署的相关信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `deploy_service_records`
--

/*!40000 ALTER TABLE `deploy_service_records` DISABLE KEYS */;
/*!40000 ALTER TABLE `deploy_service_records` ENABLE KEYS */;

--
-- Table structure for table `envs`
--

DROP TABLE IF EXISTS `envs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `envs` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键：环境记录唯一标识，自增',
  `name` varchar(255) DEFAULT NULL COMMENT '环境名称',
  `type` varchar(255) DEFAULT NULL COMMENT '环境类型',
  `project_name` varchar(255) DEFAULT NULL COMMENT '所属项目名称',
  `namespace` varchar(255) DEFAULT NULL COMMENT '环境的命名空间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`,`project_name`) COMMENT '唯一索引：环境名称和项目名称组合，确保环境名称在项目中的唯一性',
  KEY `project_name` (`project_name`) COMMENT '索引：项目名称，用于快速查找环境信息',
  CONSTRAINT `envs_ibfk_1` FOREIGN KEY (`project_name`) REFERENCES `projects` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='环境表，存储与环境相关的信息，包括环境名称、类型、所属项目和命名空间';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `envs`
--

/*!40000 ALTER TABLE `envs` DISABLE KEYS */;
/*!40000 ALTER TABLE `envs` ENABLE KEYS */;

--
-- Table structure for table `projects`
--

DROP TABLE IF EXISTS `projects`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `projects` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键：服务记录唯一标识，自增',
  `name` varchar(255) NOT NULL COMMENT '服务名称',
  `env` json DEFAULT NULL COMMENT '服务的环境配置信息，以 JSON 格式存储',
  `clusters` json DEFAULT NULL COMMENT '服务所属的集群信息，以 JSON 格式存储',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`) COMMENT '唯一索引：服务名称，确保服务名称唯一'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='服务表，存储与服务相关的信息，包括名称、环境配置和集群信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `projects`
--

/*!40000 ALTER TABLE `projects` DISABLE KEYS */;
/*!40000 ALTER TABLE `projects` ENABLE KEYS */;

--
-- Table structure for table `services`
--

DROP TABLE IF EXISTS `services`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `services` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键：服务记录唯一标识，自增',
  `name` varchar(255) DEFAULT NULL COMMENT '服务名称',
  `project_name` varchar(255) DEFAULT NULL COMMENT '服务所属项目名称',
  `env_name` varchar(255) DEFAULT NULL COMMENT '服务部署环境名称',
  `code_library_name` varchar(255) DEFAULT NULL COMMENT '代码库名称',
  `deploy_map` json DEFAULT NULL COMMENT '服务的部署映射，存储服务相关的部署信息',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`) COMMENT '唯一索引：服务名称，确保服务名称唯一',
  KEY `project_name` (`project_name`) COMMENT '索引：项目名称，用于加速查询',
  CONSTRAINT `services_ibfk_1` FOREIGN KEY (`project_name`) REFERENCES `projects` (`name`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='服务表，存储与服务相关的信息，包括名称、环境配置和集群信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `services`
--

/*!40000 ALTER TABLE `services` DISABLE KEYS */;
/*!40000 ALTER TABLE `services` ENABLE KEYS */;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户唯一标识，自增',
  `name` varchar(50) NOT NULL COMMENT '用户名，唯一',
  `email` varchar(255) DEFAULT NULL COMMENT '用户电子邮件地址',
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户电话',
  `last_login` timestamp NULL DEFAULT NULL COMMENT '用户最后登录时间',
  `token` varchar(255) DEFAULT NULL COMMENT '用户登录令牌',
  `password` varchar(255) NOT NULL COMMENT '用户密码',
  `type` varchar(255) DEFAULT NULL COMMENT '用户类型',
  `role` varchar(255) DEFAULT NULL COMMENT '用户角色',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户表，存储系统用户的相关信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'admin','admin@neolix.cn','00000000000','2025-04-28 03:42:07','YWRtaW46aTdUSUNBaUw=','$2a$10$6zsDHgyUq2/098MNsiwMw.dlVYMeWUMjyZGgCxOnfNpJ28ANDgZsC','local','admin','2025-04-27 02:37:36'),(2,'user','user@neolix.cn','00000000000','2025-04-28 03:42:07','dXNlcjpaUlB1dm9CbQ==','$2a$10$6zsDHgyUq2/098MNsiwMw.dlVYMeWUMjyZGgCxOnfNpJ28ANDgZsC','local','user','2025-04-28 03:41:10');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;

--
-- Dumping routines for database 'odyssey'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-04-28 14:43:20
