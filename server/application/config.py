import os

class Config(object):
    SECRET_KEY = os.urandom(24).hex()
    SECRET_TOKEN = os.urandom(24).hex()

class ProductionConfig(Config):
    pass

class DevelopmentConfig(Config):
    DEBUG = True

class TestingConfig(Config):
    TESTING = True