-- otus.`user` definition

CREATE TABLE `user` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `surname` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `age` int(11) NOT NULL DEFAULT '0',
  `city` varchar(100) NOT NULL,
  `interests` varchar(768) NOT NULL,
  `sex` int(11) NOT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `user_email_ukey` (`email`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4;

ALTER TABLE `user` ADD token varchar(255) NULL;
ALTER TABLE `user` CHANGE token token varchar(255) NULL AFTER password;

-- otus.user_relation definition

CREATE TABLE `user_relation` (
  `relation_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `friend_user_id` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`relation_id`),
  UNIQUE KEY `user_relation__user_id_friend_user_id_ukey` (`user_id`,`friend_user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8;
