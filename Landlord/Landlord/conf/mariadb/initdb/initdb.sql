CREATE TABLE IF NOT EXISTS history (
  Hostname VARCHAR(50),
  Username VARCHAR(50),
  Time VARCHAR(50),
  Command VARCHAR(255)
);
CREATE TABLE IF NOT EXISTS inotify (
  Hostname VARCHAR(50),
  Time VARCHAR(50),
  Event VARCHAR(20),
  Path VARCHAR(255)
);
