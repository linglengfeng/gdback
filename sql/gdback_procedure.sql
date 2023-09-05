---------------------------------------------------------------------------------------------------------------------
DROP PROCEDURE IF EXISTS `sp_user_insert`;
delimiter ;;
CREATE PROCEDURE `sp_user_insert`(in in_account varchar(128), in_password varchar(128))
BEGIN
select count(*) into @count from `user` where `account` = in_account;
if @count = 0 then
    insert into `user` (`account`, `password`) values (in_account, in_password);
    select 'r1';
else
    select 'r2';
end if;
END
;;
delimiter ;

---------------------------------------------------------------------------------------------------------------------
DROP PROCEDURE IF EXISTS `sp_user_login`;
delimiter ;;
CREATE PROCEDURE `sp_user_login`(in in_account varchar(128), in_password varchar(128))
BEGIN
select count(*) into @count from `user` where `account` = in_account and `password` = in_password;
if @count = 0 then
    select 'r2';
else
    select 'r1';
end if;
END
;;
delimiter ;
