# fulcrum

## Postgresql Setup

```bash
sudo pacman -S postgresql

sudo -iu postgres
initdb -D /var/lib/postgres/data
exit

sudo systemctl enable postgresql.service
sudo systemctl start postgresql.service

sudo -iu postgres
createuser --interactive
createdb <username>
createdb fulcrum
exit

psql -d fulcrum -f setup.sql
```

```bash
sudo apt install postgresql postgresql-contrib
sudo service postgresql start

sudo -iu postgres
createuser --interactive
createdb <username>
createdb fulcrum
exit

psql -d fulcrum -f setup.sql
```