-- phpMyAdmin SQL Dump
-- version 4.8.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1:3306
-- Creato il: Ago 23, 2018 alle 19:37
-- Versione del server: 8.0.11
-- Versione PHP: 7.2.8

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `agenda`
--
DROP DATABASE IF EXISTS `agenda`;
CREATE DATABASE IF NOT EXISTS `agenda` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
USE `agenda`;

-- --------------------------------------------------------

--
-- Struttura della tabella `npjmx_jevents_vevdetail`
--

CREATE TABLE `npjmx_jevents_vevdetail` (
  `evdet_id` int(12) NOT NULL,
  `rawdata` longtext,
  `dtstart` int(11) NOT NULL DEFAULT '0',
  `dtstartraw` varchar(30) NOT NULL DEFAULT '',
  `duration` int(11) NOT NULL DEFAULT '0',
  `durationraw` varchar(30) NOT NULL DEFAULT '',
  `dtend` int(11) NOT NULL DEFAULT '0',
  `dtendraw` varchar(30) NOT NULL DEFAULT '',
  `dtstamp` varchar(30) NOT NULL DEFAULT '',
  `class` varchar(10) NOT NULL DEFAULT '',
  `categories` varchar(120) NOT NULL DEFAULT '',
  `color` varchar(20) NOT NULL DEFAULT '',
  `description` longtext NOT NULL,
  `geolon` float NOT NULL DEFAULT '0',
  `geolat` float NOT NULL DEFAULT '0',
  `location` varchar(120) NOT NULL DEFAULT '',
  `priority` tinyint(3) NOT NULL DEFAULT '0',
  `status` varchar(20) NOT NULL DEFAULT '',
  `summary` longtext NOT NULL,
  `contact` varchar(120) NOT NULL DEFAULT '',
  `organizer` varchar(120) NOT NULL DEFAULT '',
  `url` text,
  `extra_info` text,
  `created` varchar(30) NOT NULL DEFAULT '',
  `sequence` int(11) NOT NULL DEFAULT '1',
  `state` tinyint(3) NOT NULL DEFAULT '1',
  `modified` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `multiday` tinyint(3) NOT NULL DEFAULT '1',
  `hits` int(11) NOT NULL DEFAULT '0',
  `noendtime` tinyint(3) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dump dei dati per la tabella `npjmx_jevents_vevdetail`
--

INSERT INTO `npjmx_jevents_vevdetail` (`evdet_id`, `rawdata`, `dtstart`, `dtstartraw`, `duration`, `durationraw`, `dtend`, `dtendraw`, `dtstamp`, `class`, `categories`, `color`, `description`, `geolon`, `geolat`, `location`, `priority`, `status`, `summary`, `contact`, `organizer`, `url`, `extra_info`, `created`, `sequence`, `state`, `modified`, `multiday`, `hits`, `noendtime`) VALUES
(1, NULL, 1445275631, '', 0, '', 1506235952, '', '', '', '', '', 'morbi non lectus aliquam sit amet diam in magna bibendum imperdiet nullam orci', 0, 0, '', 0, '', 'risus praesent lectus vestibulum', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:18', 1, 0, 0),
(2, NULL, 1481561914, '', 0, '', 1510948668, '', '', '', '', '', 'tortor quis turpis sed ante vivamus tortor duis mattis egestas', 0, 0, '', 0, '', 'habitasse platea dictumst etiam faucibus', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:18', 1, 0, 0),
(3, NULL, 1462511382, '', 0, '', 1516193811, '', '', '', '', '', 'tellus in sagittis dui vel nisl duis ac nibh fusce', 0, 0, '', 0, '', 'nulla eget eros elementum', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:18', 1, 0, 0),
(4, NULL, 1443003930, '', 0, '', 1515606951, '', '', '', '', '', 'tempus vivamus in felis eu sapien cursus vestibulum proin eu', 0, 0, '', 0, '', 'mi integer ac', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:18', 1, 0, 0),
(5, NULL, 1444386259, '', 0, '', 1516733187, '', '', '', '', '', 'ut mauris eget massa tempor convallis nulla neque libero convallis', 0, 0, '', 0, '', 'ac tellus', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:18', 1, 0, 0),
(6, NULL, 1477280858, '', 0, '', 1510279230, '', '', '', '', '', 'amet sem fusce consequat nulla nisl nunc nisl duis bibendum felis sed interdum venenatis turpis', 0, 0, '', 0, '', 'cubilia curae donec', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:18', 1, 0, 0),
(7, NULL, 1443248561, '', 0, '', 1509469421, '', '', '', '', '', 'in faucibus orci luctus et ultrices posuere cubilia curae', 0, 0, '', 0, '', 'nunc commodo placerat praesent', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:18', 1, 0, 0),
(8, NULL, 1488412009, '', 0, '', 1504308234, '', '', '', '', '', 'rutrum neque aenean auctor gravida sem praesent id massa id', 0, 0, '', 0, '', 'felis sed lacus morbi sem', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:19', 1, 0, 0),
(9, NULL, 1475010986, '', 0, '', 1520286695, '', '', '', '', '', 'aliquet at feugiat non pretium quis lectus suspendisse potenti in eleifend quam a', 0, 0, '', 0, '', 'in quam fringilla rhoncus mauris', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:19', 1, 0, 0),
(10, NULL, 1435290366, '', 0, '', 1518255205, '', '', '', '', '', 'in quis justo maecenas rhoncus aliquam lacus morbi quis tortor id nulla ultrices aliquet', 0, 0, '', 0, '', 'at lorem integer', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:19', 1, 0, 0),
(11, NULL, 1481283277, '', 0, '', 1495896073, '', '', '', '', '', 'lectus aliquam sit amet diam in magna bibendum imperdiet nullam', 0, 0, '', 0, '', 'curae mauris viverra diam', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:19', 1, 0, 0),
(12, NULL, 1466342903, '', 0, '', 1524686999, '', '', '', '', '', 'pede libero quis orci nullam molestie nibh in lectus', 0, 0, '', 0, '', 'ante vel ipsum praesent blandit', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:19', 1, 0, 0),
(13, NULL, 1466358391, '', 0, '', 1508632238, '', '', '', '', '', 'bibendum felis sed interdum venenatis turpis enim blandit mi in porttitor pede justo eu', 0, 0, '', 0, '', 'integer a', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:19', 1, 0, 0),
(14, NULL, 1457160356, '', 0, '', 1507231983, '', '', '', '', '', 'diam vitae quam suspendisse potenti nullam porttitor lacus at turpis donec', 0, 0, '', 0, '', 'vulputate ut ultrices vel augue', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:19', 1, 0, 0),
(15, NULL, 1450693491, '', 0, '', 1518933630, '', '', '', '', '', 'amet turpis elementum ligula vehicula consequat morbi a', 0, 0, '', 0, '', 'turpis nec euismod scelerisque', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:19', 1, 0, 0),
(16, NULL, 1453712266, '', 0, '', 1506575455, '', '', '', '', '', 'nulla elit ac nulla sed vel enim sit amet nunc viverra dapibus', 0, 0, '', 0, '', 'nascetur ridiculus mus vivamus vestibulum', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:19', 1, 0, 0),
(17, NULL, 1487026420, '', 0, '', 1495025656, '', '', '', '', '', 'sed tincidunt eu felis fusce posuere felis sed lacus morbi', 0, 0, '', 0, '', 'arcu libero rutrum ac', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:20', 1, 0, 0),
(18, NULL, 1481304366, '', 0, '', 1502129084, '', '', '', '', '', 'at vulputate vitae nisl aenean lectus pellentesque eget nunc donec quis', 0, 0, '', 0, '', 'habitasse platea', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:20', 1, 0, 0),
(19, NULL, 1433263737, '', 0, '', 1495182123, '', '', '', '', '', 'est donec odio justo sollicitudin ut suscipit a', 0, 0, '', 0, '', 'non ligula pellentesque', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:20', 1, 0, 0),
(20, NULL, 1461431990, '', 0, '', 1499972338, '', '', '', '', '', 'amet eleifend pede libero quis orci nullam molestie nibh in', 0, 0, '', 0, '', 'adipiscing lorem', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:20', 1, 0, 0),
(21, NULL, 1476349765, '', 0, '', 1507310355, '', '', '', '', '', 'commodo vulputate justo in blandit ultrices enim lorem ipsum dolor', 0, 0, '', 0, '', 'adipiscing molestie hendrerit', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:20', 1, 0, 0),
(22, NULL, 1478583141, '', 0, '', 1511671914, '', '', '', '', '', 'vivamus metus arcu adipiscing molestie hendrerit at vulputate vitae nisl aenean lectus', 0, 0, '', 0, '', 'pellentesque viverra pede ac', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:20', 1, 0, 0),
(23, NULL, 1467132122, '', 0, '', 1494177997, '', '', '', '', '', 'lacinia erat vestibulum sed magna at nunc commodo', 0, 0, '', 0, '', 'convallis tortor risus dapibus', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:20', 1, 0, 0),
(24, NULL, 1452206358, '', 0, '', 1493855205, '', '', '', '', '', 'sed interdum venenatis turpis enim blandit mi in porttitor pede justo eu massa', 0, 0, '', 0, '', 'in ante', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:20', 1, 0, 0),
(25, NULL, 1488869884, '', 0, '', 1512068108, '', '', '', '', '', 'morbi ut odio cras mi pede malesuada in imperdiet', 0, 0, '', 0, '', 'egestas metus aenean fermentum', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:20', 1, 0, 0),
(26, NULL, 1484346555, '', 0, '', 1499348832, '', '', '', '', '', 'egestas metus aenean fermentum donec ut mauris eget massa tempor', 0, 0, '', 0, '', 'molestie nibh in lectus', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:21', 1, 0, 0),
(27, NULL, 1493036640, '', 0, '', 1501137379, '', '', '', '', '', 'vitae quam suspendisse potenti nullam porttitor lacus at turpis donec posuere metus vitae', 0, 0, '', 0, '', 'vestibulum eget vulputate ut', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:21', 1, 0, 0),
(28, NULL, 1486113709, '', 0, '', 1518473806, '', '', '', '', '', 'morbi quis tortor id nulla ultrices aliquet maecenas leo odio condimentum id luctus', 0, 0, '', 0, '', 'justo sollicitudin ut suscipit a', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:21', 1, 0, 0),
(29, NULL, 1437261417, '', 0, '', 1509626442, '', '', '', '', '', 'massa quis augue luctus tincidunt nulla mollis molestie lorem quisque ut erat', 0, 0, '', 0, '', 'et tempus semper est quam', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:21', 1, 0, 0),
(30, NULL, 1461166204, '', 0, '', 1520108708, '', '', '', '', '', 'orci vehicula condimentum curabitur in libero ut massa volutpat convallis morbi odio odio elementum eu', 0, 0, '', 0, '', 'habitasse platea', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:21', 1, 0, 0),
(31, NULL, 1458680988, '', 0, '', 1499182968, '', '', '', '', '', 'ultrices aliquet maecenas leo odio condimentum id luctus nec molestie sed justo pellentesque viverra pede', 0, 0, '', 0, '', 'augue vestibulum rutrum rutrum neque', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:21', 1, 0, 0),
(32, NULL, 1453517578, '', 0, '', 1503199918, '', '', '', '', '', 'lacinia nisi venenatis tristique fusce congue diam id ornare', 0, 0, '', 0, '', 'suscipit nulla elit ac', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:21', 1, 0, 0),
(33, NULL, 1485381784, '', 0, '', 1513967398, '', '', '', '', '', 'vestibulum ante ipsum primis in faucibus orci luctus et', 0, 0, '', 0, '', 'pulvinar lobortis', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:21', 1, 0, 0),
(34, NULL, 1447521234, '', 0, '', 1493782918, '', '', '', '', '', 'aliquet pulvinar sed nisl nunc rhoncus dui vel sem sed sagittis nam congue risus semper', 0, 0, '', 0, '', 'at turpis', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:21', 1, 0, 0),
(35, NULL, 1491868767, '', 0, '', 1505243023, '', '', '', '', '', 'aliquam lacus morbi quis tortor id nulla', 0, 0, '', 0, '', 'venenatis tristique fusce', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:21', 1, 0, 0),
(36, NULL, 1481979768, '', 0, '', 1511475868, '', '', '', '', '', 'id luctus nec molestie sed justo pellentesque viverra pede ac diam cras', 0, 0, '', 0, '', 'nullam molestie nibh in lectus', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:22', 1, 0, 0),
(37, NULL, 1462055740, '', 0, '', 1498457031, '', '', '', '', '', 'tortor risus dapibus augue vel accumsan tellus nisi eu orci mauris lacinia sapien', 0, 0, '', 0, '', 'vitae nisl aenean lectus', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:22', 1, 0, 0),
(38, NULL, 1441498448, '', 0, '', 1497429723, '', '', '', '', '', 'turpis elementum ligula vehicula consequat morbi a ipsum', 0, 0, '', 0, '', 'quam suspendisse potenti nullam porttitor', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:22', 1, 0, 0),
(39, NULL, 1459861809, '', 0, '', 1507055031, '', '', '', '', '', 'posuere cubilia curae donec pharetra magna vestibulum aliquet ultrices erat tortor sollicitudin mi sit', 0, 0, '', 0, '', 'phasellus in felis', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:22', 1, 0, 0),
(40, NULL, 1454187658, '', 0, '', 1504184838, '', '', '', '', '', 'vel accumsan tellus nisi eu orci mauris lacinia sapien quis libero nullam sit amet turpis', 0, 0, '', 0, '', 'at nibh in hac', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:22', 1, 0, 0),
(41, NULL, 1448652078, '', 0, '', 1503110317, '', '', '', '', '', 'bibendum imperdiet nullam orci pede venenatis non sodales', 0, 0, '', 0, '', 'viverra eget congue eget semper', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:22', 1, 0, 0),
(42, NULL, 1483356382, '', 0, '', 1510274048, '', '', '', '', '', 'molestie nibh in lectus pellentesque at nulla suspendisse potenti cras in', 0, 0, '', 0, '', 'rhoncus mauris', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:22', 1, 0, 0),
(43, NULL, 1442767939, '', 0, '', 1504718177, '', '', '', '', '', 'sollicitudin ut suscipit a feugiat et eros vestibulum', 0, 0, '', 0, '', 'eleifend luctus ultricies eu', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:22', 1, 0, 0),
(44, NULL, 1485755950, '', 0, '', 1517873690, '', '', '', '', '', 'purus eu magna vulputate luctus cum sociis natoque penatibus et magnis dis parturient montes', 0, 0, '', 0, '', 'ipsum dolor sit amet', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:22', 1, 0, 0),
(45, NULL, 1434100010, '', 0, '', 1512409692, '', '', '', '', '', 'consequat dui nec nisi volutpat eleifend donec ut', 0, 0, '', 0, '', 'vitae mattis nibh ligula nec', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:22', 1, 0, 0),
(46, NULL, 1438144905, '', 0, '', 1495479030, '', '', '', '', '', 'faucibus orci luctus et ultrices posuere cubilia curae nulla dapibus', 0, 0, '', 0, '', 'pede libero', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:23', 1, 0, 0),
(47, NULL, 1469538792, '', 0, '', 1508420189, '', '', '', '', '', 'sem praesent id massa id nisl venenatis lacinia aenean sit amet justo', 0, 0, '', 0, '', 'ac est lacinia', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:23', 1, 0, 0),
(48, NULL, 1444287530, '', 0, '', 1522915356, '', '', '', '', '', 'vel enim sit amet nunc viverra dapibus nulla suscipit ligula in lacus curabitur', 0, 0, '', 0, '', 'massa volutpat convallis morbi', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:23', 1, 0, 0),
(49, NULL, 1448499691, '', 0, '', 1511899528, '', '', '', '', '', 'pede lobortis ligula sit amet eleifend pede libero quis orci nullam molestie nibh in lectus', 0, 0, '', 0, '', 'nunc viverra dapibus nulla suscipit', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:23', 1, 0, 0),
(50, NULL, 1434697149, '', 0, '', 1504556449, '', '', '', '', '', 'adipiscing molestie hendrerit at vulputate vitae nisl', 0, 0, '', 0, '', 'leo odio', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:23', 1, 0, 0),
(51, NULL, 1477043036, '', 0, '', 1512146887, '', '', '', '', '', 'proin at turpis a pede posuere nonummy', 0, 0, '', 0, '', 'erat eros viverra', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:23', 1, 0, 0),
(52, NULL, 1488083782, '', 0, '', 1497010555, '', '', '', '', '', 'turpis enim blandit mi in porttitor pede justo eu massa donec dapibus', 0, 0, '', 0, '', 'sit amet', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:23', 1, 0, 0),
(53, NULL, 1483815705, '', 0, '', 1507589728, '', '', '', '', '', 'non interdum in ante vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere', 0, 0, '', 0, '', 'maecenas ut', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:23', 1, 0, 0),
(54, NULL, 1492299506, '', 0, '', 1517258606, '', '', '', '', '', 'non velit nec nisi vulputate nonummy maecenas tincidunt lacus at velit vivamus vel nulla eget', 0, 0, '', 0, '', 'molestie sed justo pellentesque', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:23', 1, 0, 0),
(55, NULL, 1442155484, '', 0, '', 1498456752, '', '', '', '', '', 'luctus et ultrices posuere cubilia curae donec pharetra magna vestibulum', 0, 0, '', 0, '', 'odio cras mi pede malesuada', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:24', 1, 0, 0),
(56, NULL, 1478588425, '', 0, '', 1512470418, '', '', '', '', '', 'mi integer ac neque duis bibendum morbi non quam nec dui luctus rutrum', 0, 0, '', 0, '', 'quisque id justo sit amet', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:24', 1, 0, 0),
(57, NULL, 1434912895, '', 0, '', 1503475391, '', '', '', '', '', 'suspendisse accumsan tortor quis turpis sed ante vivamus tortor duis mattis', 0, 0, '', 0, '', 'magna ac', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:24', 1, 0, 0),
(58, NULL, 1486623055, '', 0, '', 1501226019, '', '', '', '', '', 'turpis sed ante vivamus tortor duis mattis egestas metus', 0, 0, '', 0, '', 'ultrices aliquet', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:24', 1, 0, 0),
(59, NULL, 1432257614, '', 0, '', 1519866789, '', '', '', '', '', 'curabitur at ipsum ac tellus semper interdum mauris ullamcorper purus', 0, 0, '', 0, '', 'est lacinia nisi', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:24', 1, 0, 0),
(60, NULL, 1486498599, '', 0, '', 1520351601, '', '', '', '', '', 'a feugiat et eros vestibulum ac est lacinia nisi venenatis tristique fusce congue diam', 0, 0, '', 0, '', 'posuere cubilia curae', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:24', 1, 0, 0),
(61, NULL, 1477882607, '', 0, '', 1521993084, '', '', '', '', '', 'eros suspendisse accumsan tortor quis turpis sed ante', 0, 0, '', 0, '', 'lacus morbi sem', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:24', 1, 0, 0),
(62, NULL, 1445402761, '', 0, '', 1519315505, '', '', '', '', '', 'elementum in hac habitasse platea dictumst morbi vestibulum velit id pretium iaculis diam', 0, 0, '', 0, '', 'cum sociis', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:24', 1, 0, 0),
(63, NULL, 1466524222, '', 0, '', 1513565046, '', '', '', '', '', 'porttitor lacus at turpis donec posuere metus vitae ipsum aliquam non mauris morbi non', 0, 0, '', 0, '', 'erat quisque erat', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:24', 1, 0, 0),
(64, NULL, 1449704521, '', 0, '', 1509972733, '', '', '', '', '', 'sapien placerat ante nulla justo aliquam quis turpis eget elit sodales', 0, 0, '', 0, '', 'quisque arcu libero rutrum', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:24', 1, 0, 0),
(65, NULL, 1464563272, '', 0, '', 1503446448, '', '', '', '', '', 'faucibus orci luctus et ultrices posuere cubilia curae donec pharetra magna vestibulum', 0, 0, '', 0, '', 'habitasse platea dictumst', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:25', 1, 0, 0),
(66, NULL, 1492379559, '', 0, '', 1513773480, '', '', '', '', '', 'enim sit amet nunc viverra dapibus nulla suscipit ligula in lacus curabitur at ipsum', 0, 0, '', 0, '', 'consectetuer adipiscing elit proin interdum', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:25', 1, 0, 0),
(67, NULL, 1439043815, '', 0, '', 1516008979, '', '', '', '', '', 'non mattis pulvinar nulla pede ullamcorper augue a suscipit nulla elit', 0, 0, '', 0, '', 'odio odio elementum eu interdum', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:25', 1, 0, 0),
(68, NULL, 1443412332, '', 0, '', 1511673459, '', '', '', '', '', 'in faucibus orci luctus et ultrices posuere cubilia curae mauris viverra diam', 0, 0, '', 0, '', 'ante ipsum primis in', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:25', 1, 0, 0),
(69, NULL, 1465433267, '', 0, '', 1498842795, '', '', '', '', '', 'proin at turpis a pede posuere nonummy integer non velit donec diam neque', 0, 0, '', 0, '', 'ut blandit', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:25', 1, 0, 0),
(70, NULL, 1474715261, '', 0, '', 1505440383, '', '', '', '', '', 'adipiscing elit proin risus praesent lectus vestibulum quam sapien varius ut blandit', 0, 0, '', 0, '', 'sit amet sem fusce consequat', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:25', 1, 0, 0),
(71, NULL, 1437267800, '', 0, '', 1521463778, '', '', '', '', '', 'lacus morbi sem mauris laoreet ut rhoncus aliquet pulvinar sed nisl nunc', 0, 0, '', 0, '', 'non velit nec', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:25', 1, 0, 0),
(72, NULL, 1440956528, '', 0, '', 1517542561, '', '', '', '', '', 'nam congue risus semper porta volutpat quam pede lobortis ligula sit', 0, 0, '', 0, '', 'curae nulla dapibus', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:26', 1, 0, 0),
(73, NULL, 1457916870, '', 0, '', 1510742611, '', '', '', '', '', 'proin leo odio porttitor id consequat in consequat ut nulla', 0, 0, '', 0, '', 'blandit nam', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:26', 1, 0, 0),
(74, NULL, 1483435333, '', 0, '', 1521794311, '', '', '', '', '', 'sapien urna pretium nisl ut volutpat sapien arcu sed augue', 0, 0, '', 0, '', 'et ultrices posuere cubilia curae', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:26', 1, 0, 0),
(75, NULL, 1453929224, '', 0, '', 1521445458, '', '', '', '', '', 'nisi nam ultrices libero non mattis pulvinar nulla pede ullamcorper augue a', 0, 0, '', 0, '', 'a pede', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:26', 1, 0, 0),
(76, NULL, 1463698367, '', 0, '', 1514353965, '', '', '', '', '', 'mus etiam vel augue vestibulum rutrum rutrum neque', 0, 0, '', 0, '', 'nisi eu orci mauris lacinia', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:26', 1, 0, 0),
(77, NULL, 1474864298, '', 0, '', 1495728789, '', '', '', '', '', 'sed augue aliquam erat volutpat in congue etiam', 0, 0, '', 0, '', 'ultrices posuere cubilia curae mauris', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:26', 1, 0, 0),
(78, NULL, 1476480528, '', 0, '', 1499104340, '', '', '', '', '', 'lectus pellentesque eget nunc donec quis orci eget', 0, 0, '', 0, '', 'nisi volutpat eleifend donec ut', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:26', 1, 0, 0),
(79, NULL, 1485057671, '', 0, '', 1521913958, '', '', '', '', '', 'nisi venenatis tristique fusce congue diam id ornare', 0, 0, '', 0, '', 'potenti in eleifend quam a', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:26', 1, 0, 0),
(80, NULL, 1487772075, '', 0, '', 1494652818, '', '', '', '', '', 'gravida sem praesent id massa id nisl', 0, 0, '', 0, '', 'lobortis vel dapibus', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:26', 1, 0, 0),
(81, NULL, 1447640897, '', 0, '', 1521549467, '', '', '', '', '', 'vulputate elementum nullam varius nulla facilisi cras', 0, 0, '', 0, '', 'convallis tortor risus dapibus augue', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:27', 1, 0, 0),
(82, NULL, 1485498470, '', 0, '', 1508880024, '', '', '', '', '', 'congue elementum in hac habitasse platea dictumst morbi vestibulum velit id pretium iaculis', 0, 0, '', 0, '', 'natoque penatibus et', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:27', 1, 0, 0),
(83, NULL, 1485278369, '', 0, '', 1522592428, '', '', '', '', '', 'adipiscing molestie hendrerit at vulputate vitae nisl aenean lectus', 0, 0, '', 0, '', 'cum sociis natoque penatibus et', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:27', 1, 0, 0),
(84, NULL, 1464680507, '', 0, '', 1511226296, '', '', '', '', '', 'non pretium quis lectus suspendisse potenti in eleifend quam a odio in hac habitasse', 0, 0, '', 0, '', 'eu interdum', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:27', 1, 0, 0),
(85, NULL, 1465250241, '', 0, '', 1513730575, '', '', '', '', '', 'hac habitasse platea dictumst maecenas ut massa quis augue', 0, 0, '', 0, '', 'diam vitae quam suspendisse', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:27', 1, 0, 0),
(86, NULL, 1438152753, '', 0, '', 1511418691, '', '', '', '', '', 'mi sit amet lobortis sapien sapien non mi integer ac neque', 0, 0, '', 0, '', 'justo aliquam quis turpis', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:27', 1, 0, 0),
(87, NULL, 1430963392, '', 0, '', 1523083258, '', '', '', '', '', 'nibh quisque id justo sit amet sapien dignissim vestibulum vestibulum ante ipsum primis in', 0, 0, '', 0, '', 'montes nascetur', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:27', 1, 0, 0),
(88, NULL, 1442339669, '', 0, '', 1494586932, '', '', '', '', '', 'non interdum in ante vestibulum ante ipsum primis in faucibus orci luctus et ultrices', 0, 0, '', 0, '', 'amet erat', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:27', 1, 0, 0),
(89, NULL, 1440917909, '', 0, '', 1524058467, '', '', '', '', '', 'magnis dis parturient montes nascetur ridiculus mus vivamus vestibulum sagittis sapien', 0, 0, '', 0, '', 'pretium iaculis', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:27', 1, 0, 0),
(90, NULL, 1449652715, '', 0, '', 1523287259, '', '', '', '', '', 'nulla ut erat id mauris vulputate elementum nullam varius', 0, 0, '', 0, '', 'porttitor lacus at turpis donec', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:28', 1, 0, 0),
(91, NULL, 1488509805, '', 0, '', 1504414775, '', '', '', '', '', 'ac neque duis bibendum morbi non quam nec', 0, 0, '', 0, '', 'sit amet justo', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:28', 1, 0, 0),
(92, NULL, 1473786521, '', 0, '', 1509021498, '', '', '', '', '', 'morbi vel lectus in quam fringilla rhoncus mauris enim leo', 0, 0, '', 0, '', 'fusce consequat nulla nisl nunc', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:28', 1, 0, 0),
(93, NULL, 1435647792, '', 0, '', 1515970876, '', '', '', '', '', 'ultrices posuere cubilia curae duis faucibus accumsan odio curabitur convallis duis consequat dui', 0, 0, '', 0, '', 'quam fringilla', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:28', 1, 0, 0),
(94, NULL, 1461205483, '', 0, '', 1512579436, '', '', '', '', '', 'pulvinar sed nisl nunc rhoncus dui vel sem sed sagittis nam congue risus semper porta', 0, 0, '', 0, '', 'eu interdum', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:28', 1, 0, 0),
(95, NULL, 1449825195, '', 0, '', 1508155570, '', '', '', '', '', 'quam pede lobortis ligula sit amet eleifend pede libero quis', 0, 0, '', 0, '', 'posuere cubilia curae mauris viverra', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:28', 1, 0, 0),
(96, NULL, 1432191609, '', 0, '', 1513329906, '', '', '', '', '', 'ullamcorper purus sit amet nulla quisque arcu libero rutrum ac lobortis vel', 0, 0, '', 0, '', 'vel ipsum', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:28', 1, 0, 0),
(97, NULL, 1479353482, '', 0, '', 1505292451, '', '', '', '', '', 'dui luctus rutrum nulla tellus in sagittis dui vel nisl duis ac', 0, 0, '', 0, '', 'consequat in consequat ut', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:28', 1, 0, 0),
(98, NULL, 1440646420, '', 0, '', 1498495792, '', '', '', '', '', 'faucibus orci luctus et ultrices posuere cubilia curae donec pharetra magna vestibulum aliquet ultrices erat', 0, 0, '', 0, '', 'nisi nam ultrices libero non', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:28', 1, 0, 0),
(99, NULL, 1452312283, '', 0, '', 1504669224, '', '', '', '', '', 'nulla neque libero convallis eget eleifend luctus ultricies eu', 0, 0, '', 0, '', 'aenean sit amet', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:29', 1, 0, 0),
(100, NULL, 1447977074, '', 0, '', 1524936189, '', '', '', '', '', 'faucibus accumsan odio curabitur convallis duis consequat dui nec', 0, 0, '', 0, '', 'risus dapibus augue vel accumsan', '', '', NULL, NULL, '', 1, 1, '2018-04-29 18:25:29', 1, 0, 0);

--
-- Indici per le tabelle scaricate
--

--
-- Indici per le tabelle `npjmx_jevents_vevdetail`
--
ALTER TABLE `npjmx_jevents_vevdetail`
  ADD PRIMARY KEY (`evdet_id`);

--
-- AUTO_INCREMENT per le tabelle scaricate
--

--
-- AUTO_INCREMENT per la tabella `npjmx_jevents_vevdetail`
--
ALTER TABLE `npjmx_jevents_vevdetail`
  MODIFY `evdet_id` int(12) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=101;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
