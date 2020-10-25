/*
SQLyog Community v13.1.7 (64 bit)
MySQL - 10.4.11-MariaDB : Database - mahasiswa
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`mahasiswa` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `mahasiswa`;

/*Table structure for table `mahasiswa` */

DROP TABLE IF EXISTS `mahasiswa`;

CREATE TABLE `mahasiswa` (
  `nobp` int(15) NOT NULL,
  `nama` varchar(50) NOT NULL,
  `jurusan` varchar(50) NOT NULL,
  `jalan` varchar(50) NOT NULL,
  `kelurahan` varchar(50) NOT NULL,
  `kecamatan` varchar(50) NOT NULL,
  `kabupaten` varchar(50) NOT NULL,
  `provinsi` varchar(50) NOT NULL,
  PRIMARY KEY (`nobp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `mahasiswa` */

/*Table structure for table `nilai` */

DROP TABLE IF EXISTS `nilai`;

CREATE TABLE `nilai` (
  `nobp` int(15) NOT NULL,
  `ip` varchar(50) NOT NULL,
  `semester` varchar(50) NOT NULL,
  PRIMARY KEY (`nobp`),
  CONSTRAINT `nilai_ibfk_1` FOREIGN KEY (`nobp`) REFERENCES `mahasiswa` (`nobp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `nilai` */

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
