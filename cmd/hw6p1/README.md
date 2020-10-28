# HW 6

## Results
time ./hw6p1

user75432 has access to file hello with permissions W

All users that have access to file 'world': [{user0 X} {user1000 X} {user2000 RX} {user3000 RW} {user4000 WX} {user5000 R} {user6000 RX} {user7000 W} {user8000 W} {user9000 X} {user10000 WX} {user11000 X} {user12000 R} {user13000 X} {user14000 WX} {user15000 X} {user16000 WX} {user17000 WX} {user18000 RX} {user19000 WX} {user20000 W} {user21000 X} {user22000 R} {user23000 X} {user24000 WX} {user25000 W} {user26000 WX} {user27000 R} {user28000 RW} {user29000 RW} {user30000 RX} {user31000 WX} {user32000 WX} {user33000 RW} {user34000 RX} {user35000 RX} {user36000 RW} {user37000 RX} {user38000 WX} {user39000 RW} {user40000 RX} {user41000 RX} {user42000 WX} {user43000 WX} {user44000 WX} {user45000 X} {user46000 RX} {user47000 RX} {user48000 W} {user49000 R} {user50000 W} {user51000 RX} {user52000 W} {user53000 RW} {user54000 W} {user55000 RX} {user56000 X} {user57000 RX} {user58000 W} {user59000 WX} {user60000 WX} {user61000 X} {user62000 WX} {user63000 RX} {user64000 X} {user65000 RX} {user66000 RW} {user67000 W} {user68000 WX} {user69000 RX} {user70000 RW} {user71000 X} {user72000 W} {user73000 WX} {user74000 R} {user75000 R} {user76000 RX} {user77000 RW} {user78000 W} {user79000 W} {user80000 W} {user81000 RX} {user82000 W} {user83000 RX} {user84000 RX} {user85000 WX} {user86000 RX} {user87000 RX} {user88000 RW} {user89000 RX} {user90000 WX} {user91000 R} {user92000 RX} {user93000 RW} {user94000 WX} {user95000 X} {user96000 WX} {user97000 W} {user98000 X} {user99000 WX}]

./hw6p1  0.05s user 0.05s system 103% cpu 0.090 total

As you can see, the results are pretty quick for adding 100,000 users and then doing one query by username and one query by filename (with a file that has 100 associated users with some sort of permissions attached.)