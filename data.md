Build a basic messaging communication system with a producer and consumer app. 

The producer will have a basic CRUD method and on CREATE, UPDATE and DELETE events, send the payload through the queue and consume by the Consumer, and sink into a Postgress Database.


Producer App:

    implement CRUD method for handling data
    use message queue like channels to send messages on CREATE, UPDATE and DELETE

Consumer App:
    listen to the message queue
    Process incoming messages


Postgres : store and manage


