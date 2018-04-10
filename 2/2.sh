mysql -u root -pqwerty123 test -e "SELECT u.user, (SELECT COUNT(*) FROM messages WHERE UID = u.UID) TotalMessages FROM users u"
