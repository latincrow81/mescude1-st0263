class RmqSubscriberService:

    def __init__(self):
        self.config = {
            'host': 'localhost',
            'port': 5672,
            'exchange': 'topic'
        }