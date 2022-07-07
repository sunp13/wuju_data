CREATE DATABASE  IF NOT EXISTS `b365api` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin */;
USE `b365api`;
-- MySQL dump 10.13  Distrib 8.0.26, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: b365api
-- ------------------------------------------------------
-- Server version	5.7.34

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `asian_handicap`
--

DROP TABLE IF EXISTS `asian_handicap`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `asian_handicap` (
  `hand_id` int(11) NOT NULL AUTO_INCREMENT,
  `comm_id` varchar(45) COLLATE utf8mb4_bin DEFAULT '' COMMENT '对应比赛id\n',
  `home_odds` varchar(45) COLLATE utf8mb4_bin DEFAULT '',
  `home_handicap` varchar(45) COLLATE utf8mb4_bin DEFAULT '',
  `away_odds` varchar(45) COLLATE utf8mb4_bin DEFAULT '' COMMENT '更新时间\n',
  `away_handicap` varchar(45) COLLATE utf8mb4_bin DEFAULT '',
  `update_at` int(11) DEFAULT '0',
  `is_deleted` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`hand_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='亚盘';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `asian_handicap`
--

LOCK TABLES `asian_handicap` WRITE;
/*!40000 ALTER TABLE `asian_handicap` DISABLE KEYS */;
/*!40000 ALTER TABLE `asian_handicap` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `full_time`
--

DROP TABLE IF EXISTS `full_time`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `full_time` (
  `full_id` int(11) NOT NULL AUTO_INCREMENT,
  `comm_id` varchar(45) COLLATE utf8mb4_bin DEFAULT '',
  `home_odds` varchar(45) COLLATE utf8mb4_bin DEFAULT '',
  `draw_odds` varchar(45) COLLATE utf8mb4_bin DEFAULT '',
  `away_odds` varchar(45) COLLATE utf8mb4_bin DEFAULT '',
  `update_at` int(11) DEFAULT '0',
  `is_deleted` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`full_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='欧';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `full_time`
--

LOCK TABLES `full_time` WRITE;
/*!40000 ALTER TABLE `full_time` DISABLE KEYS */;
/*!40000 ALTER TABLE `full_time` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goal_line`
--

DROP TABLE IF EXISTS `goal_line`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `goal_line` (
  `goal_id` int(11) NOT NULL AUTO_INCREMENT,
  `comm_id` varchar(45) COLLATE utf8mb4_bin DEFAULT '',
  `over_odds` varchar(45) COLLATE utf8mb4_bin DEFAULT '',
  `over_name` varchar(45) COLLATE utf8mb4_bin DEFAULT '',
  `under_odds` varchar(45) COLLATE utf8mb4_bin DEFAULT '',
  `under_name` varchar(45) COLLATE utf8mb4_bin DEFAULT '',
  `update_at` int(11) DEFAULT '0',
  `is_deleted` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`goal_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='大小';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goal_line`
--

LOCK TABLES `goal_line` WRITE;
/*!40000 ALTER TABLE `goal_line` DISABLE KEYS */;
/*!40000 ALTER TABLE `goal_line` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `up_coming`
--

DROP TABLE IF EXISTS `up_coming`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `up_coming` (
  `comm_id` varchar(45) COLLATE utf8mb4_bin NOT NULL,
  `comm_time` int(11) NOT NULL,
  `league_id` varchar(45) COLLATE utf8mb4_bin DEFAULT '' COMMENT '联赛ID',
  `league_name` varchar(200) COLLATE utf8mb4_bin DEFAULT '' COMMENT '联赛名称',
  `home_id` varchar(45) COLLATE utf8mb4_bin DEFAULT '' COMMENT '主队ID',
  `home_name` varchar(200) COLLATE utf8mb4_bin DEFAULT '' COMMENT '主队名称',
  `away_id` varchar(45) COLLATE utf8mb4_bin DEFAULT '' COMMENT '客队ID',
  `away_name` varchar(45) COLLATE utf8mb4_bin DEFAULT '' COMMENT '客队名称',
  `ss` varchar(45) COLLATE utf8mb4_bin DEFAULT '',
  `our_event_id` varchar(45) COLLATE utf8mb4_bin DEFAULT '',
  `r_id` varchar(45) COLLATE utf8mb4_bin DEFAULT '',
  `update_at` varchar(45) COLLATE utf8mb4_bin DEFAULT '',
  `is_deleted` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`comm_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='即将比赛';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `up_coming`
--

LOCK TABLES `up_coming` WRITE;
/*!40000 ALTER TABLE `up_coming` DISABLE KEYS */;
/*!40000 ALTER TABLE `up_coming` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-07-03  1:50:04
