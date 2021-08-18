-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Aug 18, 2021 at 10:36 PM
-- Server version: 10.4.17-MariaDB
-- PHP Version: 7.3.27

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `movie_db`
--

-- --------------------------------------------------------

--
-- Table structure for table `link_urls`
--

CREATE TABLE `link_urls` (
  `id_trailer` bigint(20) NOT NULL,
  `url` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `log_wishlist`
--

CREATE TABLE `log_wishlist` (
  `id_log_wishlist` int(11) NOT NULL,
  `movie_id` int(11) NOT NULL,
  `movie_title` varchar(255) NOT NULL,
  `movie_year` varchar(50) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id_user` bigint(20) NOT NULL,
  `first_name` longtext DEFAULT NULL,
  `last_name` longtext DEFAULT NULL,
  `user_name` longtext DEFAULT NULL,
  `email` longtext DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `wishlish`
--

CREATE TABLE `wishlish` (
  `id_wishlist` int(11) NOT NULL,
  `id_log_wishlist` int(11) NOT NULL,
  `id_user` bigint(20) NOT NULL,
  `movie_id` int(11) NOT NULL,
  `movie_title` varchar(255) NOT NULL,
  `movie_year` varchar(45) NOT NULL,
  `movie_rated` varchar(45) NOT NULL,
  `movie_released` varchar(100) NOT NULL,
  `movie_runtime` varchar(100) NOT NULL,
  `movie_genre` varchar(255) NOT NULL,
  `movie_director` varchar(255) NOT NULL,
  `movie_writer` varchar(255) NOT NULL,
  `movie_actor` varchar(255) NOT NULL,
  `movie_plot` text NOT NULL,
  `movie_language` varchar(255) NOT NULL,
  `movie_country` varchar(255) NOT NULL,
  `movie_award` varchar(255) NOT NULL,
  `movie_poster` varchar(255) NOT NULL,
  `movie_imdb_rating` varchar(255) NOT NULL,
  `movie_is_completed` int(11) NOT NULL,
  `created_at` datetime NOT NULL,
  `modified_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `link_urls`
--
ALTER TABLE `link_urls`
  ADD PRIMARY KEY (`id_trailer`);

--
-- Indexes for table `log_wishlist`
--
ALTER TABLE `log_wishlist`
  ADD PRIMARY KEY (`id_log_wishlist`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id_user`);

--
-- Indexes for table `wishlish`
--
ALTER TABLE `wishlish`
  ADD PRIMARY KEY (`id_wishlist`),
  ADD KEY `id_user` (`id_user`),
  ADD KEY `movie_id` (`movie_id`),
  ADD KEY `id_log_wishlist` (`id_log_wishlist`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `link_urls`
--
ALTER TABLE `link_urls`
  MODIFY `id_trailer` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `log_wishlist`
--
ALTER TABLE `log_wishlist`
  MODIFY `id_log_wishlist` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id_user` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `wishlish`
--
ALTER TABLE `wishlish`
  MODIFY `id_wishlist` int(11) NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `wishlish`
--
ALTER TABLE `wishlish`
  ADD CONSTRAINT `wishlish_ibfk_1` FOREIGN KEY (`id_user`) REFERENCES `users` (`id_user`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `wishlish_ibfk_2` FOREIGN KEY (`id_log_wishlist`) REFERENCES `log_wishlist` (`id_log_wishlist`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
