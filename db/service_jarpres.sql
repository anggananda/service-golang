-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 14 Des 2024 pada 12.53
-- Versi server: 10.4.27-MariaDB
-- Versi PHP: 8.2.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `service_jarpres`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `inventories`
--

CREATE TABLE `inventories` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `description` text DEFAULT NULL,
  `quantity` int(11) NOT NULL DEFAULT 0,
  `status` enum('available','reserved','damaged') DEFAULT 'available',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `inventories`
--

INSERT INTO `inventories` (`id`, `name`, `description`, `quantity`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(3, 'Laptop Dell XPS 13', 'Ultrabook high-performance', 10, 'available', '2024-12-14 04:59:39', '2024-12-14 09:01:00', NULL),
(4, 'Monitor Samsung 27 Inch', 'High-resolution display', 30, 'available', '2024-12-14 04:59:39', '2024-12-14 08:17:58', NULL),
(5, 'Keyboard Mechanical Razer', 'RGB backlit mechanical keyboard', 50, 'available', '2024-12-14 05:00:40', '2024-12-14 08:18:01', NULL),
(6, 'Mouse Logitech MX Master 3', 'Ergonomic wireless mouse', 25, 'available', '2024-12-14 05:00:40', '2024-12-14 08:18:03', NULL),
(7, 'External Hard Drive 1TB', 'Portable storage device', 40, 'reserved', '2024-12-14 05:01:00', '2024-12-14 09:01:06', NULL),
(8, 'test', 'test', 30, 'available', '2024-12-14 07:40:36', '2024-12-14 08:18:08', NULL);

-- --------------------------------------------------------

--
-- Struktur dari tabel `tasks`
--

CREATE TABLE `tasks` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `vehicle_id` int(11) DEFAULT NULL,
  `inventory_id` int(11) DEFAULT NULL,
  `destination` varchar(255) NOT NULL,
  `status` enum('pending','in_progress','completed','cancelled') DEFAULT 'pending',
  `assigned_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `completed_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `tasks`
--

INSERT INTO `tasks` (`id`, `user_id`, `vehicle_id`, `inventory_id`, `destination`, `status`, `assigned_at`, `completed_at`, `created_at`, `updated_at`, `deleted_at`) VALUES
(2, 1, 7, 7, 'karangasem', 'completed', '2024-12-14 11:06:24', NULL, '2024-12-14 11:06:24', '2024-12-14 11:47:45', NULL),
(3, 1, 7, 7, 'Siangaraja', 'in_progress', '2024-12-14 11:47:40', NULL, '2024-12-14 11:47:40', '2024-12-14 11:47:40', NULL),
(4, 1, 5, 3, 'Denpasar', 'pending', '2024-12-14 11:48:17', NULL, '2024-12-14 11:48:17', '2024-12-14 11:48:17', NULL),
(8, 1, 4, 4, 'Bangli', 'pending', '2024-12-14 11:49:33', NULL, '2024-12-14 11:49:33', '2024-12-14 11:49:33', NULL);

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `role` enum('admin','staff','driver') NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `role`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'angga', 'anggadek857@gmail.com', 'driver', '2024-12-14 08:07:35', '2024-12-14 11:20:00', NULL);

-- --------------------------------------------------------

--
-- Struktur dari tabel `vehicles`
--

CREATE TABLE `vehicles` (
  `id` int(11) NOT NULL,
  `license_plate` varchar(50) NOT NULL,
  `brand` varchar(100) NOT NULL,
  `status` enum('active','inactive','maintenance') DEFAULT 'active',
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `vehicles`
--

INSERT INTO `vehicles` (`id`, `license_plate`, `brand`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(4, 'DK 1234 AB', 'Toyota', 'active', '2024-12-14 05:04:25', '2024-12-14 09:01:16', NULL),
(5, 'DK 5678 CD', 'Honda', 'active', '2024-12-14 05:04:25', '2024-12-14 09:01:19', NULL),
(6, 'DK 9101 EF', 'Suzuki', 'active', '2024-12-14 05:04:25', '2024-12-14 05:04:25', NULL),
(7, 'DK 1213 GH', 'Mitsubishi', 'active', '2024-12-14 05:04:25', '2024-12-14 09:01:24', NULL),
(8, 'DK 1415 IJ', 'Ford', 'maintenance', '2024-12-14 05:04:25', '2024-12-14 09:01:46', NULL),
(13, 'DK 1919 SS', 'Honda', 'active', '2024-12-14 07:24:00', '2024-12-14 07:24:00', NULL),
(14, 'DK 8888 SD', 'Nissan', 'inactive', '2024-12-14 10:10:02', '2024-12-14 10:10:02', NULL);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `inventories`
--
ALTER TABLE `inventories`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `tasks`
--
ALTER TABLE `tasks`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `vehicle_id` (`vehicle_id`),
  ADD KEY `inventory_id` (`inventory_id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Indeks untuk tabel `vehicles`
--
ALTER TABLE `vehicles`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `license_plate` (`license_plate`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `inventories`
--
ALTER TABLE `inventories`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT untuk tabel `tasks`
--
ALTER TABLE `tasks`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT untuk tabel `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `vehicles`
--
ALTER TABLE `vehicles`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `tasks`
--
ALTER TABLE `tasks`
  ADD CONSTRAINT `tasks_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `tasks_ibfk_2` FOREIGN KEY (`vehicle_id`) REFERENCES `vehicles` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `tasks_ibfk_3` FOREIGN KEY (`inventory_id`) REFERENCES `inventories` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
