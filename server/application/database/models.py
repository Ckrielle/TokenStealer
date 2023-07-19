from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import Column, Integer, String, create_engine

engine = create_engine("sqlite:///storage.sqlite3", echo=False)

Base = declarative_base()

class DiscordUser(Base):
    __tablename__ = 'discordusers'
    id = Column(Integer, primary_key=True)
    discord_token = Column(String)
    discord_username = Column(String)
    email = Column(String)
    phone = Column(String)

Base.metadata.create_all(engine)