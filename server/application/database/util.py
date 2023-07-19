from application.database.models import DiscordUser
from sqlalchemy.orm import sessionmaker
from sqlalchemy import create_engine

engine = create_engine("sqlite:///storage.sqlite3", echo=False)
Session = sessionmaker(bind=engine)
session = Session()

def addDiscordUser(token, handle, email, phone):
    new_user = DiscordUser(
        discord_token=token,
        discord_username=handle,
        email=email,
        phone=phone
    )
    session.add(new_user)
    session.commit()

def getAllDiscordUsers():
    return session.query(DiscordUser).all()