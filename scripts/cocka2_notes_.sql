-- phpMyAdmin SQL Dump
-- version 4.9.5deb2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Feb 03, 2021 at 05:14 PM
-- Server version: 8.0.23-0ubuntu0.20.04.1
-- PHP Version: 7.4.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `cocka2_notes`
--
CREATE DATABASE IF NOT EXISTS `cocka2_notes` DEFAULT CHARACTER SET utf8mb4;
USE `cocka2_notes`;

-- --------------------------------------------------------

--
-- Table structure for table `checkbox_note_item`
--

CREATE TABLE `checkbox_note_item` (
  `id` bigint NOT NULL,
  `text` varchar(150) DEFAULT NULL,
  `checked` tinyint NOT NULL DEFAULT '0',
  `note_id` bigint NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `checkbox_note_item`
--

INSERT INTO `checkbox_note_item` (`id`, `text`, `checked`, `note_id`) VALUES
(1, 'bread', 1, 1),
(2, 'all spice', 0, 1),
(6, 'Cheese sticks', 0, 1),
(7, 'Spinich', 1, 1),
(10, 'peppers', 1, 1),
(11, 'Eggs', 0, 1),
(19, 'milk', 0, 6),
(20, 'dog food', 0, 1),
(23, 'xxx', 1, 9),
(24, 'ccc', 1, 9),
(26, 'ss', 0, 10),
(27, 'wheat bread', 0, 6),
(28, 'spinish', 0, 6),
(30, 'chicken breast', 0, 6),
(31, 'sugar', 0, 6),
(35, 'Milk', 0, 1),
(36, 'testf', 0, 3),
(37, '', 0, 3),
(38, '', 0, 7),
(39, 'brown sugar', 0, 6),
(40, '', 0, 7),
(41, '', 0, 7),
(42, '', 0, 7),
(43, '', 0, 7),
(44, 'ggggg', 0, 6),
(45, '', 0, 7),
(46, '', 0, 7),
(47, '', 0, 7),
(48, '', 0, 7),
(49, 'coffees', 0, 6),
(53, 'creamer', 0, 6),
(54, 'chips', 0, 6),
(55, 'Shoes', 1, 12),
(56, 'Shirt', 1, 12),
(62, 'test', 0, 13),
(65, 'Bread', 0, 18),
(66, 'Milk', 0, 18),
(67, 'good stuff', 0, 16),
(68, 'good stuff', 0, 16),
(69, 'this store is closed', 0, 19),
(71, 'good stuff', 0, 19),
(74, 'good stuff', 0, 13),
(75, 'ken', 0, 13),
(77, 'ddddddd', 0, 22),
(78, 'ffffff', 0, 22);

-- --------------------------------------------------------

--
-- Table structure for table `mail_server`
--

CREATE TABLE `mail_server` (
  `id` bigint NOT NULL,
  `mail_server` varchar(150) NOT NULL,
  `secure_connection` tinyint NOT NULL DEFAULT '0',
  `port` int NOT NULL DEFAULT '25',
  `debug` tinyint NOT NULL DEFAULT '0',
  `username` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `send_address` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `mail_server`
--

INSERT INTO `mail_server` (`id`, `mail_server`, `secure_connection`, `port`, `debug`, `username`, `password`, `send_address`) VALUES
(1, 'smtpout.secureserver.net', 0, 465, 0, 'support@myapigateway.com', 'V0RlYzQyMTg=', 'support@myapigateway.com');

-- --------------------------------------------------------

--
-- Table structure for table `note`
--

CREATE TABLE `note` (
  `id` bigint NOT NULL,
  `title` varchar(45) NOT NULL,
  `type` varchar(10) NOT NULL DEFAULT 'checkbox',
  `owner_email` varchar(75) NOT NULL,
  `last_used` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `note`
--

INSERT INTO `note` (`id`, `title`, `type`, `owner_email`, `last_used`) VALUES
(1, 'Sams Club 1/5/2021', 'checkbox', 'tester@tester.com', '2021-01-25 22:00:33'),
(3, 'some text note with no checkbox only text', 'note', 'tester@tester.com', '2021-01-25 23:27:51'),
(4, 'a note to end all notes and then some', 'checkbox', 'test@test.com', '2021-01-04 15:38:46'),
(6, 'Kroger', 'checkbox', 'tester@tester.com', '2021-01-26 15:05:42'),
(7, 'Home Depot', 'note', 'tester@tester.com', '2021-01-26 15:59:56'),
(8, 'rrrrr', 'checkbox', 'tester1@tester.com', '2021-01-17 22:10:23'),
(9, 'aaaaa', 'checkbox', 'tester1@tester.com', '2021-01-17 22:10:49'),
(10, 'ddddd', 'checkbox', 'tester1@tester.com', '2021-01-17 22:14:27'),
(11, 'notes', 'note', 'tester1@tester.com', '2021-01-17 23:23:50'),
(12, 'Academy sprorts', 'checkbox', 'tester@tester.com', '2021-01-26 23:21:11'),
(13, 'Costco', 'checkbox', 'tester@tester.com', '2021-01-26 23:21:37'),
(16, 'another note 1/27/2021', 'checkbox', 'tester@tester.com', '2021-01-27 16:57:14'),
(18, 'Sams Club', 'checkbox', 'kenwwilliamson12@gmail.com', '2021-01-27 22:33:26'),
(19, 'Echards', 'checkbox', 'tester@tester.com', '2021-01-28 15:59:40'),
(22, 'aaaaa', 'checkbox', 'tester@tester.com', '2021-02-03 18:54:34');

-- --------------------------------------------------------

--
-- Table structure for table `note_item`
--

CREATE TABLE `note_item` (
  `id` bigint NOT NULL,
  `text` varchar(500) DEFAULT NULL,
  `note_id` bigint NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `note_item`
--

INSERT INTO `note_item` (`id`, `text`, `note_id`) VALUES
(2, 'some note', 3),
(5, 'Do some stuff ff', 3),
(6, 'do some more stuff', 3),
(7, 'Emp# : 44577777-999', 7),
(10, 'Phone: 458-555-8484', 7),
(15, 'aaaa', 11),
(16, 'bbbbb', 11),
(17, 'Email: ken@ken.com', 7),
(18, 'good stuff', 7),
(19, 'milk', 7),
(20, 'good stuff-99', 7),
(24, 'fbbccd', 7),
(27, 'bread', 7),
(28, 'ddddd', 7),
(34, 'good stuff', 7),
(36, 'test some', 3),
(37, 'hhhhh', 3),
(38, 'funny stuff', 7),
(40, 'hello', 3),
(41, 'test', 7),
(42, 'ssn : *** *** 3456', 3),
(61, 'bbbbb', 7),
(70, '', 3),
(88, 'aaaaa', 7);

-- --------------------------------------------------------

--
-- Table structure for table `note_users`
--

CREATE TABLE `note_users` (
  `user_email` varchar(75) NOT NULL,
  `note_id` bigint NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `note_users`
--

INSERT INTO `note_users` (`user_email`, `note_id`) VALUES
('tester@tester.com', 1),
('tester@tester.com', 3),
('test@test.com', 4),
('tester@tester.com', 6),
('tester@tester.com', 7),
('tester1@tester.com', 8),
('tester1@tester.com', 9),
('tester1@tester.com', 10),
('tester1@tester.com', 11),
('test@test.com', 7),
('tester@tester.com', 12),
('tester@tester.com', 13),
('tester@tester.com', 16),
('kenwwilliamson12@gmail.com', 18),
('tester@tester.com', 19),
('tester@tester.com', 22);

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `email` varchar(75) NOT NULL,
  `password` varchar(1000) NOT NULL,
  `web_enabled` tinyint NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`email`, `password`, `web_enabled`) VALUES
('', '', 0),
('kenwwilliamson12@gmail.com', '$2a$10$gx6Y6d7mGfYCyFiSw59u3OyOXUwGL6nrXJ.3bk7ZRwuz3MveNa/Fi', 0),
('test@test.com', '$2a$10$GBC2IuYvscpE.wXRLjFL0.r8MYdbSwfqMblrXaNDiT.TOyplDtmvO', 0),
('tester@test.com', '$2a$10$5UgxP4yeFmMjNESTbfp7bObyWkXXvU7OyYgROZzPFbTw.5Oh6PBo2', 0),
('tester@tester.com', '$2a$10$xEhc2AUlY5H.4nYfa56u/.y4VhvqRC5ahU66VBG44bDiCJukkokhe', 0),
('tester1@tester.com', '', 0),
('tester2@tester.com', '$2a$10$OAyRPLkchujrE/GahEVgfu.zkcrtKTE2UgJa7HedGZXPU6EzCp6XC', 0);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `checkbox_note_item`
--
ALTER TABLE `checkbox_note_item`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id_UNIQUE` (`id`),
  ADD KEY `fk_checkbox_note_item_note1_idx` (`note_id`);

--
-- Indexes for table `mail_server`
--
ALTER TABLE `mail_server`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `note`
--
ALTER TABLE `note`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id_UNIQUE` (`id`),
  ADD KEY `fk_note_user1_idx` (`owner_email`);

--
-- Indexes for table `note_item`
--
ALTER TABLE `note_item`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id_UNIQUE` (`id`),
  ADD KEY `fk_note_item_note1_idx` (`note_id`);

--
-- Indexes for table `note_users`
--
ALTER TABLE `note_users`
  ADD KEY `fk_note_users_user1_idx` (`user_email`),
  ADD KEY `fk_note_users_note1_idx` (`note_id`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`email`),
  ADD UNIQUE KEY `email_UNIQUE` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `checkbox_note_item`
--
ALTER TABLE `checkbox_note_item`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=79;

--
-- AUTO_INCREMENT for table `mail_server`
--
ALTER TABLE `mail_server`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `note`
--
ALTER TABLE `note`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=24;

--
-- AUTO_INCREMENT for table `note_item`
--
ALTER TABLE `note_item`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=89;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `checkbox_note_item`
--
ALTER TABLE `checkbox_note_item`
  ADD CONSTRAINT `fk_checkbox_note_item_note1` FOREIGN KEY (`note_id`) REFERENCES `note` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `note`
--
ALTER TABLE `note`
  ADD CONSTRAINT `fk_note_user1` FOREIGN KEY (`owner_email`) REFERENCES `user` (`email`) ON DELETE CASCADE;

--
-- Constraints for table `note_item`
--
ALTER TABLE `note_item`
  ADD CONSTRAINT `fk_note_item_note1` FOREIGN KEY (`note_id`) REFERENCES `note` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `note_users`
--
ALTER TABLE `note_users`
  ADD CONSTRAINT `fk_note_users_note1` FOREIGN KEY (`note_id`) REFERENCES `note` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `fk_note_users_user1` FOREIGN KEY (`user_email`) REFERENCES `user` (`email`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
