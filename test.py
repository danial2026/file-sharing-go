from kafka import KafkaProducer
import json

producer = KafkaProducer(bootstrap_servers='localhost:9092')

#producer.send('GroupID', b'some_message_bytes')
r = {'is_claimed': 'True', 'rating': 3.5}
r = json.dumps(r)

producer.send('message-log', str.encode(r))
