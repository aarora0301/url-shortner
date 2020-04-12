#### Cassandra setup

1. [Setup](https://www.javatpoint.com/how-to-install-cassandra-on-mac) cassandra on mac .
2. Start Cassandra server: `` Cassandra -f ``
3. Enter into the shell mode with command: `` cqlsh``
4. Run following queries:

````
a. CREATE KEYSPACE = dev
     WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 }; 
     
b. CREATE TABLE url (
       hash text PRIMARY KEY,
       creation_date timestamp,
       expiration_date timestamp,
       original_url text,
       user_id int
   ) 
   
c. CREATE TABLE used_key (
       key text PRIMARY KEY
   ) 
   
d. CREATE TABLE available_key (
       key text PRIMARY KEY
   ) 
   
  