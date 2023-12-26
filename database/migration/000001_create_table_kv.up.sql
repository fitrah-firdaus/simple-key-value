CREATE TABLE IF NOT EXISTS `kv` (
  `keylog` varchar(255) NOT NULL,
  `value` text NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
ALTER TABLE `kv`
  ADD PRIMARY KEY (`keylog`);
