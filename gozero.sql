-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jan 18, 2024 at 04:58 AM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.0.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `gozero`
--

-- --------------------------------------------------------

--
-- Table structure for table `articles`
--

CREATE TABLE `articles` (
  `article_id` int(11) NOT NULL,
  `title` varchar(150) NOT NULL,
  `image` varchar(50) NOT NULL,
  `text` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `articles`
--

INSERT INTO `articles` (`article_id`, `title`, `image`, `text`) VALUES
(1, 'Tips Memilih Skincare Ramah Lingkungan Ala Maurilla Imron', '', 'Banyak yang belum tahu kalau gaya hidup minimalis sebenarnya bisa diterapkan di skincare. Zaman sekarang, banyak yang menganggap untuk menjaga kulit wajib memakai berbagai printilan dari skincare atau melakukannya dengan rentetan step yang panjang. Hal ini tentu akan menghasilkan semakin banyak sampah bekas produk skincare karena demand market yang tinggi. Belum lagi sangat banyak brand skincare yang tidak mempedulikan lingkungan saat memproduksi dan menjual produk mereka. Walaupun sudah banyak yang aware tentang skincare ramah lingkungan, namun tidak banyak yang tahu bagaimana cara memilih skincare ramah lingkungan yang tepat. Apalagi sekarang ini sudah sangat banyak jenis dan brand skincare di pasaran. '),
(2, 'Manfaat Kompos Untuk Tanaman', '', 'Diketahui dari salah satu blog Waste4change yang menyatakan, sisa makanan masih mendominasi di Tempat Pemrosesan Akhir (TPA), bahkan jumlahnya jauh lebih banyak dibanding sampah plastik, yakni sebesar 28,3 persen dari 21,53 juta ton. Padahal, sampah organik adalah sampah yang paling mudah untuk terurai dan memiliki banyak manfaat jika dikelola dengan baik, salah satunya dikompos.\r\n<br><br>\r\nKompos adalah bahan organik yang dihasilkan dari proses penguraian bahan organik seperti dedaunan, ranting, sampah dapur, dan kotoran hewan. Proses pengomposan dapat dilakukan secara alami dengan membiarkan bahan organik terurai secara alami oleh mikroba atau dengan cara mempercepat proses penguraian menggunakan starter kompos.\r\n<br><br>\r\nKompos memiliki manfaat yang sangat baik untuk pertumbuhan dan kesehatan tanaman. Pada artikel ini kita tidak hanya membahas mengenai manfaat, kita juga akan membahas kandungan dan cara membuat pupuk kompos.');

-- --------------------------------------------------------

--
-- Table structure for table `complete_missions`
--

CREATE TABLE `complete_missions` (
  `user_id` int(11) NOT NULL,
  `mission_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `daily_missions`
--

CREATE TABLE `daily_missions` (
  `mission_id` int(11) NOT NULL,
  `description` varchar(200) NOT NULL,
  `reward` int(11) NOT NULL,
  `link` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `daily_missions`
--

INSERT INTO `daily_missions` (`mission_id`, `description`, `reward`, `link`) VALUES
(1, 'Baca artikel', 100, '/article'),
(2, 'Kerjakan Quiz', 300, '/article'),
(3, 'Bagikan pengalamanmu di Community Gallery', 200, '/gallery');

-- --------------------------------------------------------

--
-- Table structure for table `photos`
--

CREATE TABLE `photos` (
  `photo_id` int(11) NOT NULL,
  `title` varchar(100) NOT NULL,
  `image` varchar(100) NOT NULL,
  `description` text NOT NULL,
  `user_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `quizzes`
--

CREATE TABLE `quizzes` (
  `quiz_id` int(11) NOT NULL,
  `question` text NOT NULL,
  `opt1` text NOT NULL,
  `opt2` text NOT NULL,
  `opt3` text NOT NULL,
  `answer` char(1) NOT NULL,
  `article_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `quizzes`
--

INSERT INTO `quizzes` (`quiz_id`, `question`, `opt1`, `opt2`, `opt3`, `answer`, `article_id`) VALUES
(3, 'Apa yang disarankan Maurilla Imron untuk memilih skincare ramah lingkungan yang pertama?', 'Memilih skincare lokal dengan bahan-bahan lokal', 'Memilih skincare dari brand internasional', 'Memilih skincare dengan kemasan plastik sekali pakai', '1', 1),
(4, 'Apa manfaat aplikasi INCIDecoder dalam pemilihan skincare?', 'Menyediakan tips memasak dengan minyak', 'Memberikan informasi tentang skincare yang tidak ramah lingkungan', 'Membantu mengidentifikasi ingredients dalam skincare', '3', 1),
(5, 'Apa yang harus dihindari dalam pemilihan skincare berdasarkan informasi yang diberikan?', 'Skincare dengan ingredients biodegradable', 'Skincare dengan ingredients yang berakhiran cone', 'Skincare dengan kandungan microbits', '2', 1),
(6, 'Berdasarkan artikel tersebut, bahan organik mana yang mengandung nitrogen paling tinggi?', 'Sisa sayuran dan buah', 'Ranting dan daun', 'Kotoran hewan', '1', 2),
(7, 'Berdasarkan artikel tersebut, manfaat utama penggunaan kompos untuk tanaman adalah...', 'Meningkatkan kualitas tanah', 'Meningkatkan pertumbuhan tanaman', 'Bebas penyakit, hama dan gulma', '1', 2),
(8, 'Berdasarkan artikel tersebut, cara membuat kompos yang paling mudah adalah...', 'Menggunakan starter kompos', 'Mencampur bahan organik dengan tanah', 'Menempatkan bahan organik di dalam wadah tertutup', '2', 2);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `user_id` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(50) NOT NULL,
  `point` int(11) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `articles`
--
ALTER TABLE `articles`
  ADD PRIMARY KEY (`article_id`);

--
-- Indexes for table `complete_missions`
--
ALTER TABLE `complete_missions`
  ADD KEY `user_id` (`user_id`,`mission_id`),
  ADD KEY `mission_id` (`mission_id`);

--
-- Indexes for table `daily_missions`
--
ALTER TABLE `daily_missions`
  ADD PRIMARY KEY (`mission_id`);

--
-- Indexes for table `photos`
--
ALTER TABLE `photos`
  ADD PRIMARY KEY (`photo_id`),
  ADD KEY `author_id` (`user_id`);

--
-- Indexes for table `quizzes`
--
ALTER TABLE `quizzes`
  ADD PRIMARY KEY (`quiz_id`),
  ADD KEY `article_id` (`article_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`user_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `articles`
--
ALTER TABLE `articles`
  MODIFY `article_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `daily_missions`
--
ALTER TABLE `daily_missions`
  MODIFY `mission_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `photos`
--
ALTER TABLE `photos`
  MODIFY `photo_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT for table `quizzes`
--
ALTER TABLE `quizzes`
  MODIFY `quiz_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `user_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `complete_missions`
--
ALTER TABLE `complete_missions`
  ADD CONSTRAINT `complete_missions_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `complete_missions_ibfk_2` FOREIGN KEY (`mission_id`) REFERENCES `daily_missions` (`mission_id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `photos`
--
ALTER TABLE `photos`
  ADD CONSTRAINT `photos_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `quizzes`
--
ALTER TABLE `quizzes`
  ADD CONSTRAINT `quizzes_ibfk_1` FOREIGN KEY (`article_id`) REFERENCES `articles` (`article_id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
